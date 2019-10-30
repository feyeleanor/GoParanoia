package main

import "crypto/rand"
import "crypto/rsa"
import "fmt"
import "os"

func main() {
	k, e := LoadPEM(RSA_PUBLIC_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PUBLIC_KEY)

	var b []byte
	b, e = Encrypt(k.(*rsa.PublicKey), os.Args[2])
	ExitOnError(e, RSA_ENCRYPTION_FAILED)
	fmt.Println(EncodeToString(b))
}

func Encrypt(k *rsa.PublicKey, m string) (b []byte, e error) {
	return rsa.EncryptPKCS1v15(rand.Reader, k, []byte(m))
}
