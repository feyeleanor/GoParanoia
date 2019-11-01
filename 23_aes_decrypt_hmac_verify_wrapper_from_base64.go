package main

import "fmt"
import "os"

func main() {
	s := os.Args[1]
	hs := read_base64(s[0:88])
	m := read_base64(s[88:])
	_, e := AESDecrypt(os.Getenv("AES_KEY"), m)
	ExitOnError(e, AES_DECRYPTION_FAILED)

	if HMAC_Verify(os.Getenv("HMAC_KEY"), hs, m) {
		fmt.Println("Signature Verification Succeeded")
	} else {
		fmt.Println("Signature Verification Failed")
		os.Exit(VERIFICATION_FAILED)
	}
}
