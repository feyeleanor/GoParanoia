package main

import "crypto/hmac"
import "crypto/sha512"
import "fmt"
import "os"

const (
	_ = iota
	ENCRYPTION_FAILED
)

func main() {
	var e error
	var m []byte

	k := os.Getenv("AES_KEY")
	s := os.Args[1]

	m, e = AESEncrypt(k, s)
	ExitOnError(e, ENCRYPTION_FAILED)

	hk := os.Getenv("HMAC_KEY")
	h := hmac.New(sha512.New, []byte(hk))
	h.Write(m)

	fmt.Println(
		EncodeToString(h.Sum(nil)) +
			EncodeToString([]byte(m)))
}
