package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

const (
	_ = iota
	DECRYPTION_FAILED
	INVALID_SIGNATURE
)

func main() {
	var e error

	hk := os.Getenv("HMAC_KEY")
	h := hmac.New(sha512.New, []byte(hk))

	s := os.Args[1]
	hs := []byte(read_base64(s[0:88]))
	ms := read_base64(s[88:])

	k := os.Getenv("AES_KEY")
	ms, e = AESDecrypt(k, ms)
	ExitOnError(e, DECRYPTION_FAILED)

	h.Write([]byte(ms))
	if hmac.Equal(h.Sum(nil), hs) {
		fmt.Println("Signature Verification Succeeded")
	} else {
		fmt.Println("Signature Verification Failed")
		os.Exit(INVALID_SIGNATURE)
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
