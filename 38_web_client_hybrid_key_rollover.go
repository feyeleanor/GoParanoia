package main

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
	ALICE.ShowCurrentKeys(c)

	for _, m := range os.Args[3:] {
		ALICE.Report("wants to say:", m)
		ALICE.Report("heard:", SendMessage(c, a, n, m))

		c = ChangeSymmetricKey(c, AES_MakeKey(32), a, n)
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
	r, e := http.Get(HTTP_URL(a, KEY, n))
	ExitOnError(e, WEB_REQUEST_FAILED)

	var k interface{}
	k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, HTTP_readbody(r.Body), "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, c *AES_channel, a, n string) string {
	if k := []byte(c.ki); k[len(k)-1] == 0 {
		ALICE.Report("Found one!")
		ALICE.Report("SendSymmetricKey:", k)
	}

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

func SendReceive(c *AES_channel, m string, f func(string) *http.Response) string {
	return c.DecryptMessage(
		HTTP_readbody(
			f(c.EncryptMessage(m)).Body))
}

func ChangeSymmetricKey(c *AES_channel, k, a, n string) (r *AES_channel) {
	if k := []byte(k); k[len(k)-1] == 0 {
		ALICE.Report("Found one!")
		ALICE.Report("ChangeSymmetricKey:", k)
	}
	r = &AES_channel{ko: c.ko, ki: k}
	r.ko = SendReceive(r, k, func(m string) *http.Response {
		res, e := HTTP_put(HTTP_URL(a, KEY, n), m)
		ExitOnError(e, WEB_REQUEST_FAILED)
		return res
	})
	return
}

func SendMessage(c *AES_channel, a, n, m string) string {
	me := c.EncryptMessage(m)
	r, e := HTTP_put(HTTP_URL(a, MESSAGE, n), me)
	ExitOnError(e, WEB_REQUEST_FAILED)
	return c.DecryptMessage(
		HTTP_readbody(r.Body))
}
