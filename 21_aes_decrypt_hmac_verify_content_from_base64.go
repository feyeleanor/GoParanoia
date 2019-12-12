package main

import "fmt"
import "os"

func main() {
	s := os.Args[1]
	hs := read_base64(s[0:88])
	ms, e := AES_DecryptAndTrim(
		os.Getenv("AES_KEY"),
		read_base64(s[88:]))
	ExitOnError(e, AES_DECRYPTION_FAILED)

	if HMAC_Verify(os.Getenv("HMAC_KEY"), hs, ms) {
		fmt.Println("Signature Verification Succeeded")
	} else {
		fmt.Println("Signature Verification Failed")
		os.Exit(VERIFICATION_FAILED)
	}
}
