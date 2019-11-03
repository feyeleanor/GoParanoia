package main

import "crypto/rand"
import "crypto/rsa"
import "fmt"
import "os"

func main() {
	f := os.Args[1]
	k, e := PEM_Load(RSA_PRIVATE_KEY, f, os.Getenv("PEM_KEY"))
	ExitOnError(e, INVALID_PRIVATE_KEY)

	var s string
	m := read_base64(os.Args[2])
	s, e = Decrypt(k.(*rsa.PrivateKey), m)
	ExitOnError(e, RSA_DECRYPTION_FAILED)
	fmt.Println(s)
}

func Decrypt(k *rsa.PrivateKey, m string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptPKCS1v15(rand.Reader, k, []byte(m))
	r = string(b)
	return
}
