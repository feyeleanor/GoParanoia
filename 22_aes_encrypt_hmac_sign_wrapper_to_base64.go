package main

import "fmt"
import "os"

func main() {
	m, e := AES_Encrypt(os.Getenv("AES_KEY"), os.Args[1])
	h := HMAC_Sign(os.Getenv("HMAC_KEY"), string(m))
	ExitOnError(e, AES_ENCRYPTION_FAILED)
	fmt.Println(EncodeStrings(h, m))
}
