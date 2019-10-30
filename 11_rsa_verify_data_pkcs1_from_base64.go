package main

import "crypto"
import "crypto/rsa"
import "crypto/sha512"
import "fmt"
import "os"

func main() {
	k, e := LoadPEM(RSA_PUBLIC_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PUBLIC_KEY)

	e = Verify(
		k.(*rsa.PublicKey),
		read_base64(os.Args[3]),
		os.Args[2])
	ExitOnError(e, VERIFICATION_FAILED)
	fmt.Println("Signature Verification Succeeded")
}

func Verify(k *rsa.PublicKey, s, m string) error {
	h := sha512.Sum512([]byte(m))
	return rsa.VerifyPKCS1v15(
		k,
		crypto.SHA512,
		h[:],
		[]byte(s))
}
