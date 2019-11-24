package main

import "crypto/rand"
import "crypto/rsa"
import "net/http"
import "os"

const KEY = "key"
const MESSAGE = "message"
const ALICE Person = "Alice"

func main() {
	a := ServerAddress(HTTP_ADDRESS)
	n := os.Args[1]
	c := NewSession(os.Args[2], a, n)
	ALICE.Report("received symmetric key:", EncodeToBase64(c.ko))

	for _, m := range os.Args[3:] {
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

	r, e := http.Get(HTTP_URL(a, KEY, n))
	ExitOnError(e, WEB_REQUEST_FAILED)

	k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, HTTP_readbody(r.Body), "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, c *channel, a, n string) string {
	b, e := OAEP_Encrypt(pk, c.ki, n)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)

	var r *http.Response
	r, e = http.Post(
		HTTP_URL(a, KEY, n),
		OCTET_STREAM,
		EncodeToReader(b))

	ExitOnError(e, WEB_REQUEST_FAILED)
	return c.DecryptMessage(HTTP_readbody(r.Body))
}

func SendMessage(c *channel, a, n, m string) string {
	r, e := HTTP_put(HTTP_URL(a, MESSAGE, n), c.EncryptMessage(m))
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
