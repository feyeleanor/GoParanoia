package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"

const KEY = "key"
const MESSAGE = "message"
const BOB Person = "Bob"

func init() {
	sessions := make(map[string]*AES_channel)
	k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

	priv := k.(*rsa.PrivateKey)
	p := PEM_Create(priv.PublicKey)
	HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			BOB.Report("received request for public key from", r.RemoteAddr)
			fmt.Fprint(w,
				EncodeToString(
					pem.EncodeToMemory(p)))

		case http.MethodPost:
			if n := SubPath(KEY, r); len(n) == 0 {
				http.Error(w, "missing nonce", 500)
			} else {
				s := HTTP_readbody(r.Body)
				if s, e = OAEP_Decrypt(priv, read_base64(s), n); e != nil {
					http.Error(w, "decryption failed", 500)
				} else {
					c := &AES_channel{ko: s, ki: AES_MakeKey(16)}
					sessions[n] = c
					BOB.ShowCurrentKeys(c)
					fmt.Fprint(w, c.EncryptMessage(c.ki))
				}
			}

		case http.MethodPut:
			if n := SubPath(KEY, r); len(n) == 0 {
				http.Error(w, "missing nonce", 500)
			} else if s := sessions[n]; s == nil {
				http.Error(w, "unknown nonce", 500)
			} else if m := HTTP_readbody(r.Body); m == "" {
				http.Error(w, "missing symmetric key", 500)
			} else {
				s.ko = s.DecryptMessage(m)
				s.ki = AES_MakeKey(16)
				BOB.ShowCurrentKeys(s)
				fmt.Fprint(w, s.EncryptMessage(sessions[n].ki))
			}
		}
	})

	HandleFunc(MESSAGE, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			n := SubPath(MESSAGE, r)
			if s := sessions[n]; s == nil {
				http.Error(w, "unknown nonce", 500)
			} else if m := HTTP_readbody(r.Body); m == "" {
				http.Error(w, "missing message", 500)
			} else {
				m = s.DecryptMessage(m)
				BOB.Report("heard:", m)

				m = fmt.Sprintf("%v received", m)
				BOB.Report("wants to say:", m)
				fmt.Fprint(w, s.EncryptMessage(m))
			}
		}
	})
}

func main() {
	http.ListenAndServe(
    ServerAddress(HTTP_ADDRESS), nil)
}
