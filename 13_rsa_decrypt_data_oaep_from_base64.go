package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "crypto/x509"
import "encoding/pem"
import "fmt"
import "hash"
import "io/ioutil"
import "os"

const (
	_ = iota
	MISSING_FILENAME
	FILE_UNREADABLE
	NOT_A_PEM_FILE
	NOT_A_PRIVATE_KEY
	PEM_PASSWORD_REQUIRED
	PEM_DECRYPTION_FAILED
	INVALID_PRIVATE_KEY
	DECRYPTION_FAILED
)

func main() {
	var e error
	var k *rsa.PrivateKey
	var s string

	f := os.Args[1]
	l := os.Args[2]
	m := read_base64(os.Args[3])
	b := LoadFile(f)
	p := DecodePEM(b)
	b = DecryptPEM(p)

	k, e = x509.ParsePKCS1PrivateKey(b)
	ExitOnError(e, INVALID_PRIVATE_KEY)

	s, e = Decrypt(sha512.New(), k, m, l)
	ExitOnError(e, DECRYPTION_FAILED)
	fmt.Println(s)
}

func Decrypt(h hash.Hash, k *rsa.PrivateKey, m, l string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptOAEP(h, rand.Reader, k, []byte(m), []byte(l))
	r = string(b)
	return
}

func LoadFile(s string) (b []byte) {
	var e error
	if s == "" {
		os.Exit(MISSING_FILENAME)
	}

	b, e = ioutil.ReadFile(s)
	ExitOnError(e, FILE_UNREADABLE)
	return
}

func DecodePEM(b []byte) (p *pem.Block) {
	switch p, _ = pem.Decode(b); {
	case p == nil:
		os.Exit(NOT_A_PEM_FILE)
	case p.Type != "RSA PRIVATE KEY":
		os.Exit(NOT_A_PRIVATE_KEY)
	}
	return
}

func DecryptPEM(p *pem.Block) (b []byte) {
	if x509.IsEncryptedPEMBlock(p) {
		if pwd := os.Getenv("PEM_KEY"); pwd != "" {
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
