package main

import "crypto/rsa"
import "io"
import "net/http"
import "os"
import "strings"

const KEY = "key"
const MESSAGE = "message"
const ALICE Person = "Alice"

func main() {
	a := ServerAddress(HTTP_ADDRESS)
	n := os.Args[1]
	for _, m := range os.Args[3:] {
		AsSession(os.Args[2], a, n, func(c *AES_channel) {
			ALICE.Report("using new session")
			ALICE.ShowCurrentKeys(c)
			ALICE.Report("wants to say:", m)
			ALICE.Report("heard:", SendMessage(c, a, n, m))

			ALICE.Report("changing session keys")
			c = UpdateSymmetricKey(c, a, n)
			ALICE.ShowCurrentKeys(c)
			ALICE.Report("wants to say:", m)
			ALICE.Report("heard:", SendMessage(c, a, n, m))
		})
	}
}

func NewSession(ki, a, n string) (c *AES_channel) {
	p := RequestPublicKey(a, n)
	c = &AES_channel{ki: ki}
	c.ko = SendSymmetricKey(p, c, a, n)
	return
}

func AsSession(ki, a, n string, f func(*AES_channel)) {
	c := NewSession(ki, a, n)
	defer CloseSession(c, a, n)
	f(c)
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
	r, e := http.Get(HTTP_URL(a, KEY, n))
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
		HTTP_URL(a, KEY, n),
		OCTET_STREAM,
		EncodeToReader(b))

	ExitOnError(e, WEB_REQUEST_FAILED)
	return c.DecryptMessage(HTTP_readbody(r.Body))
}

func SendReceive(c *AES_channel, m string, f func(string) *http.Response) string {
	return c.DecryptMessage(
		HTTP_readbody(
      func() io.ReadCloser {
        ALICE.Report(c.EncryptMessage(m))
  			return f(c.EncryptMessage(m)).Body
      }()))
}

func UpdateSymmetricKey(c *AES_channel, a, n string) (r *AES_channel) {
	r = &AES_channel{ki: AES_MakeKey(len(c.ki))}
	ALICE.Report("r.ki:", EncodeToBase64(r.ki))
	r.ko = SendReceive(c, r.ki, func(m string) *http.Response {
		r, e := HTTP_put(HTTP_URL(a, KEY, n), m)
		ExitOnError(e, WEB_REQUEST_FAILED)
		return r
	})
	ALICE.Report("r.ko:", EncodeToBase64(r.ko))
  if len(r.ko) % 16 != 0 {
    ALICE.Report("key received too short:", len(r.ko), "bytes")
  }
	return
}

func SendMessage(c *AES_channel, a, n, m string) string {
	return SendReceive(c, m, func(m string) *http.Response {
		r, e := HTTP_put(HTTP_URL(a, MESSAGE, n), m)
		ExitOnError(e, WEB_REQUEST_FAILED)
		return r
	})
}

func CloseSession(c *AES_channel, a, n string) {
	_, e := Delete(HTTP_URL(a, KEY, n), c.EncryptMessage(c.ki))
	ExitOnError(e, WEB_REQUEST_FAILED)
}

func Delete(url, m string) (*http.Response, error) {
	r, e := http.NewRequest("DELETE", url, strings.NewReader(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	r.ContentLength = int64(len(m))
	return http.DefaultClient.Do(r)
}
