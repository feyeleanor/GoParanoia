package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "fmt"
import "hash"
import "os"

func main() {
	k, e := PEM_Load(RSA_PUBLIC_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PUBLIC_KEY)

	l := os.Args[2]
	m := os.Args[3]
	var b []byte
	b, e = Encrypt(sha512.New(), k.(*rsa.PublicKey), m, l)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)
	fmt.Println(EncodeToString(b))
}

func Encrypt(h hash.Hash, k *rsa.PublicKey, m, l string) (b []byte, e error) {
	return rsa.EncryptOAEP(h, rand.Reader, k, []byte(m), []byte(l))
}
