package main

import "crypto/hmac"
import "crypto/sha512"
import "fmt"
import "os"

func main() {
	h := Sign(os.Getenv("HMAC_KEY"), string(os.Args[1]))
	m, e := AES_Encrypt(os.Getenv("AES_KEY"), os.Args[1])
	ExitOnError(e, AES_ENCRYPTION_FAILED)
	fmt.Println(EncodeStrings(h, m))
}

func Sign(k, m string) []byte {
	h := hmac.New(sha512.New, []byte(k))
	h.Write([]byte(m))
	return h.Sum(nil)
}
