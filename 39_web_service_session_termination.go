package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"
import "strings"
import "time"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const KEY = "/key/"
const MESSAGE = "/message/"
const OCTETS = "application/octet-stream"

const HTTP_ADDRESS = "HTTP_ADDRESS"

func init() {
	sessions := make(map[string]*AES_channel)
	k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

	priv := k.(*rsa.PrivateKey)
	p := PEM_Create(priv.PublicKey)
	http.HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			BOB.Report("received request for public key from", r.RemoteAddr)
			fmt.Fprint(w,
				EncodeToString(
					pem.EncodeToMemory(p)))

		case http.MethodPost:
			n := r.URL.Path[len(KEY):]
			if len(n) == 0 {
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
			n := r.URL.Path[len(KEY):]
			if len(n) == 0 {
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

		case http.MethodDelete:
			n := r.URL.Path[len(KEY):]
			BOB.Report("is forgotting all about:", n)
			delete(sessions, n)
		}
	})

	http.HandleFunc(MESSAGE, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			n := r.URL.Path[len(MESSAGE):]
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
	a := os.Getenv(HTTP_ADDRESS)
	if a == "" {
		a = DEFAULT_ADDRESS
	}

	go func() {
		http.ListenAndServe(a, nil)
	}()
	time.Sleep(2 * time.Second)

	n := os.Args[2]
	c := NewSession(os.Args[3], a, n)
	ALICE.ShowCurrentKeys(c)

	for _, m := range os.Args[4:] {
		ALICE.Report("wants to say:", m)
		ALICE.Report("heard:", SendMessage(c, a, n, m))

		c.ki = AES_MakeKey(16)
		c.ko = UpdateSymmetricKey(c, a, n)
		ALICE.ShowCurrentKeys(c)
	}
}

func NewSession(ki, a, n string) (c *AES_channel) {
	p := RequestPublicKey(a, n)
	ALICE.Report("received public key:", p)
	c = &AES_channel{ki: ki}
	c.ko = SendSymmetricKey(p, c, a, n)
	return
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
	r, e := http.Get(HTTP + a + KEY + n)
	ExitOnError(e, WEB_REQUEST_FAILED)

	var k interface{}
	k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, HTTP_readbody(r.Body), "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, c *AES_channel, a, n string) (s string) {
	b, e := OAEP_Encrypt(pk, c.ki, n)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)

	var r *http.Response
	r, e = http.Post(
		HTTP+a+KEY+n,
		OCTETS,
		EncodeToReader(b))

	ExitOnError(e, WEB_REQUEST_FAILED)
	return c.DecryptMessage(HTTP_readbody(r.Body))
}

func SendReceive(c *AES_channel, m string, f func(string) *http.Response) string {
	return c.DecryptMessage(
		HTTP_readbody(
			f(c.EncryptMessage(m)).Body))
}

func UpdateSymmetricKey(c *AES_channel, a, n string) (s string) {
	return SendReceive(c, c.ki, func(m string) *http.Response {
		r, e := HTTP_put(HTTP+a+KEY+n, m)
		ExitOnError(e, WEB_REQUEST_FAILED)
		return r
	})
}

func SendMessage(c *AES_channel, a, n, m string) string {
	return SendReceive(c, m, func(m string) *http.Response {
		r, e := HTTP_put(HTTP+a+MESSAGE+n, m)
		ExitOnError(e, WEB_REQUEST_FAILED)
		return r
	})
}

func Delete(url, m string) (*http.Response, error) {
	r, e := http.NewRequest("DELETE", url, strings.NewReader(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	r.ContentLength = int64(len(m))
	return http.DefaultClient.Do(r)
}
