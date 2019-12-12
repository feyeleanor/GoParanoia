package main

import "crypto/hmac"
import "fmt"
import "os"

func main() {
	s := os.Args[2]
	hs := read_base64(s[0:88])
	ms, e := AES_DecryptAndTrim(
		os.Getenv("AES_KEY"),
		read_base64(s[88:]))
	ExitOnError(e, AES_DECRYPTION_FAILED)

	switch {
	case ms != os.Args[1]:
		fmt.Println("error: content doesn't match")
		os.Exit(CONTENT_MISMATCH)
	case !Verify(os.Getenv("HMAC_KEY"), hs, os.Args[1]):
		fmt.Println("Signature Verification Failed")
		os.Exit(VERIFICATION_FAILED)
	default:
		fmt.Println("Signature Verification Succeeded")
	}
}

func Verify(k, hs, m string) bool {
	return hmac.Equal(
		[]byte(hs),
		HMAC_Sign(k, m))
}
