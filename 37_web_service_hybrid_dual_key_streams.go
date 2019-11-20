package main

import "crypto/rand"
import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"
import "time"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const KEY = "/key/"
const MESSAGE = "/message/"
const OCTETS = "application/octet-stream"

const HTTP_ADDRESS = "HTTP_ADDRESS"

func init() {
	sessions := make(map[string]*channel)
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
					c := &channel{ko: s, ki: MakeNewKey(16)}
					sessions[n] = c
					BOB.Report("received symmetric key:", s)
					BOB.Report("sends symmetric key:", EncodeToBase64(c.ki))
					fmt.Fprint(w, c.EncryptMessage(c.ki))
				}
			}
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
	ALICE.Report("received symmetric key:", EncodeToBase64(c.ko))

	for _, m := range os.Args[4:] {
		ALICE.Report("wants to say:", m)
		ALICE.Report("heard:", SendMessage(c, a, n, m))
	}
}

func NewSession(ki, a, n string) (c *channel) {
	p := RequestPublicKey(a, n)
	ALICE.Report("received public key:", p)
	c = &channel{ki: ki}
	ALICE.Report("sends symmetric key:", c.ki)
	c.ko = SendSymmetricKey(p, c, a, n)
	return
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
	var k interface{}

	r, e := http.Get(HTTP + a + KEY + n)
	ExitOnError(e, WEB_REQUEST_FAILED)

	k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, HTTP_readbody(r.Body), "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, c *channel, a, n string) (s string) {
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

func SendMessage(c *channel, a, n, m string) string {
	r, e := HTTP_put(HTTP+a+MESSAGE+n, c.EncryptMessage(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	return c.DecryptMessage(HTTP_readbody(r.Body))
}

func MakeNewKey(n int) string {
	b := make([]byte, n)
	_, e := rand.Read(b)
	ExitOnError(e, NOT_ENOUGH_RANDOMNESS)
	return string(b)
}

type channel struct{ ko, ki string }

func (a *channel) EncryptMessage(m string) string {
	b, _ := AES_Encrypt(a.ko, m)
	return EncodeToString(b)
}

func (a *channel) DecryptMessage(m string) (r string) {
	r = read_base64(m)
	r, _ = AES_Decrypt(a.ki, r)
	return
}
