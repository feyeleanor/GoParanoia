package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

const (
	_ = iota
	DECRYPTION_FAILED
	INCORRECT_CONTENT
	INVALID_SIGNATURE
)

func main() {
	var e error

	hk := os.Getenv("HMAC_KEY")
	h := hmac.New(sha512.New, []byte(hk))

	m := os.Args[1]
	h.Write([]byte(m))

	s := os.Args[2]
	hs := []byte(read_base64(s[0:88]))
	ms := read_base64(s[88:])

	k := os.Getenv("AES_KEY")
	ms, e = AESDecrypt(k, ms)
	ExitOnError(e, DECRYPTION_FAILED)

	switch {
	case ms != m:
		fmt.Println("error: content doesn't match")
		os.Exit(INCORRECT_CONTENT)
	case !hmac.Equal(h.Sum(nil), hs):
		fmt.Println("Signature Verification Failed")
		os.Exit(INVALID_SIGNATURE)
	default:
		fmt.Println("Signature Verification Succeeded")
	}
}

func read_base64(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func ExitOnError(e error, n int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(n)
	}
}
