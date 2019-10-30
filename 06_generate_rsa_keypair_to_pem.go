package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/pem"
import "io/ioutil"
import "os"
import "strconv"

const DEFAULT_KEYSIZE = 1024

func main() {
	var f string
	var n uint64
	var e error
	var k *rsa.PrivateKey

	if f = os.Args[1]; f == "" {
		os.Exit(MISSING_FILENAME)
	}
	n, e = strconv.ParseUint(os.Args[2], 10, 64)
	ExitOnError(e, INVALID_KEYSIZE)

	if n == 0 {
		n = DEFAULT_KEYSIZE
	}
	k, e = CreatePrivateKey(int(n))
	ExitOnError(e, CREATE_KEY_FAILED)

	p := CreatePEM(k)
	if pwd := os.Getenv("PEM_KEY"); pwd != "" {
		p, e = EncryptPEM(p, pwd)
		ExitOnError(e, PEM_ENCRYPTION_FAILED)
	}
	e = SaveKey(f, p, 0644)
	ExitOnError(e, FILE_WRITE_FAILED)
}

func SaveKey(f string, p *pem.Block, perm os.FileMode) error {
	return ioutil.WriteFile(f, pem.EncodeToMemory(p), perm)
}

func CreatePrivateKey(n int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, n)
}

func CreatePEM(k interface{}) (r *pem.Block) {
	switch k := k.(type) {
	case *rsa.PrivateKey:
		r = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}
	case rsa.PublicKey:
		r = &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&k),
		}
	}
	return
}

func EncryptPEM(p *pem.Block, s string) (*pem.Block, error) {
	return x509.EncryptPEMBlock(
		rand.Reader,
		p.Type,
		p.Bytes,
		[]byte(s),
		x509.PEMCipherAES256,
	)
}
