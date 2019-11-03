package main

import "crypto/rsa"
import "crypto/x509"
import "encoding/pem"
import "os"

func main() {
	f := os.Args[1]
	k, e := PEM_Load(RSA_PRIVATE_KEY, f, os.Getenv("PEM_KEY"))
	ExitOnError(e, INVALID_PRIVATE_KEY)

	f = f + ".pub"
	e = PEM_SaveKey(f, PEM_Create(k.(*rsa.PrivateKey).PublicKey), 0644)
	ExitOnError(e, FILE_WRITE_FAILED)
}

func LoadPEM(t int, f, pwd string) (k interface{}, e error) {
	b := LoadFile(f)
	p := DecodePEM(b, t)
	b = DecryptPEM(p, pwd)
	switch t {
	case RSA_PRIVATE_KEY:
		k, e = x509.ParsePKCS1PrivateKey(b)
	case RSA_PUBLIC_KEY:
		k, e = x509.ParsePKCS1PublicKey(b)
	}
	return
}

func DecodePEM(b []byte, t int) (p *pem.Block) {
	switch p, _ = pem.Decode(b); {
	case p == nil:
		os.Exit(NOT_A_PEM_FILE)
	case t == RSA_PRIVATE_KEY && p.Type != "RSA PRIVATE KEY":
		os.Exit(NOT_A_PRIVATE_KEY)
	case t == RSA_PUBLIC_KEY && p.Type != "RSA PUBLIC KEY":
		os.Exit(NOT_A_PUBLIC_KEY)
	}
	return
}

func DecryptPEM(p *pem.Block, pwd string) (b []byte) {
	if x509.IsEncryptedPEMBlock(p) {
		if pwd != "" {
			var e error
			b, e = x509.DecryptPEMBlock(p, []byte(pwd))
			ExitOnError(e, PEM_DECRYPTION_FAILED)
		} else {
			os.Exit(PEM_PASSWORD_REQUIRED)
		}
	} else {
		b = p.Bytes
	}
	return
}
