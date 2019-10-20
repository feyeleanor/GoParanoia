package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
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
	h.Write([]byte(s))

	fmt.Println(EncodeToString(h.Sum(nil), m))
}

func EncodeToString(b ...[]byte) (r string) {
	for _, v := range b {
		r += base64.StdEncoding.EncodeToString(v)
	}
	return
}

func ExitOnError(e error, n int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(n)
	}
}
