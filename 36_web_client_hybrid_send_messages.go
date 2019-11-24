package main

import "crypto/rsa"
import "net/http"
import "os"
import "strings"

const KEY = "key"
const MESSAGE = "message"
const ALICE Person = "Alice"

func main() {
	a := ServerAddress(HTTP_ADDRESS)
	n := os.Args[1]
	k := os.Args[2]
	StartSession(k, a, n)
	for _, m := range os.Args[3:] {
		ALICE.Report("wants to say:", m)
		ALICE.Report("heard:", SendMessage(a, n, k, m))
	}
}

func StartSession(k, a, n string) {
	p := RequestPublicKey(a, n)
	ALICE.Report("received public key:", p)
	ALICE.Report("sends symmetric key:", k)
	SendSymmetricKey(p, a, k, n)
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
	var k interface{}
	var s string

	r, e := http.Get(HTTP_URL(a, KEY, n))
	ExitOnError(e, WEB_REQUEST_FAILED)

	if s = HTTP_readbody(r.Body); s == "" {
		os.Exit(WEB_NO_BODY)
	}

	k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, s, "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, a, k, n string) {
	b, e := OAEP_Encrypt(pk, k, n)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)

	_, e = http.Post(
		HTTP_URL(a, KEY, n),
		OCTET_STREAM,
		EncodeToReader(b))

	ExitOnError(e, WEB_REQUEST_FAILED)
}

func SendMessage(a, n, k, m string) string {
	r, e := Put(HTTP_URL(a, MESSAGE, n), EncryptMessage(k, m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	return DecryptMessage(k, HTTP_readbody(r.Body))
}

func DecryptMessage(k, v string) string {
	v = read_base64(v)
	r, e := AES_Decrypt(k, v)
	ExitOnError(e, AES_DECRYPTION_FAILED)
	return r
}

func EncryptMessage(k, v string) string {
	b, e := AES_Encrypt(k, v)
	ExitOnError(e, AES_ENCRYPTION_FAILED)
	return EncodeToString(b)
}

func Put(url, m string) (*http.Response, error) {
	r, e := http.NewRequest("PUT", url, strings.NewReader(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	r.ContentLength = int64(len(m))
	return http.DefaultClient.Do(r)
}
