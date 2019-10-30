package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "fmt"
import "hash"
import "os"

func main() {
	k, e := LoadPEM(RSA_PRIVATE_KEY, os.Args[1], os.Getenv("PEM_KEY"))
	ExitOnError(e, INVALID_PRIVATE_KEY)

	l := os.Args[2]
	m := read_base64(os.Args[3])
	var s string
	s, e = Decrypt(sha512.New(), k.(*rsa.PrivateKey), m, l)
	ExitOnError(e, RSA_DECRYPTION_FAILED)
	fmt.Println(s)
}

func Decrypt(h hash.Hash, k *rsa.PrivateKey, m, l string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptOAEP(h, rand.Reader, k, []byte(m), []byte(l))
	r = string(b)
	return
}
