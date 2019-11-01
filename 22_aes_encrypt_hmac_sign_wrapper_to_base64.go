package main

import "fmt"
import "os"

func main() {
	m := os.Args[1]
	h := HMAC_Sign(os.Getenv("HMAC_KEY"), m)
	ms, e := AESEncrypt(os.Getenv("AES_KEY"), m)
	ExitOnError(e, AES_ENCRYPTION_FAILED)
	fmt.Println(EncodeStrings(h, ms))
}
