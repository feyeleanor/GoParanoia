package main

import "crypto/hmac"
import "crypto/sha512"
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
	h.Write([]byte(ms))

	k := os.Getenv("AES_KEY")
	ms, e = AESDecrypt(k, ms)
	ExitOnError(e, DECRYPTION_FAILED)

	if hmac.Equal(h.Sum(nil), hs) {
		fmt.Println("Signature Verification Succeeded")
	} else {
		fmt.Println("Signature Verification Failed")
		os.Exit(INVALID_SIGNATURE)
	}
}
