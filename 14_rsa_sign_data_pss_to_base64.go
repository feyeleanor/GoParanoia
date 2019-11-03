package main

import "crypto"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "fmt"
import "os"

func main() {
	k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], os.Getenv("PEM_KEY"))
	ExitOnError(e, INVALID_PRIVATE_KEY)

	var b []byte
	b, e = Sign(k.(*rsa.PrivateKey), os.Args[2], nil)
	ExitOnError(e, SIGNING_FAILED)
	fmt.Println(EncodeToString(b))
}

func Sign(k *rsa.PrivateKey, m string, o *rsa.PSSOptions) (r []byte, e error) {
	hs := sha512.Sum512([]byte(m))
	return rsa.SignPSS(rand.Reader, k, crypto.SHA512, hs[:], o)
}
