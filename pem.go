package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/pem"
import "io/ioutil"
import "os"

const (
	RSA_PRIVATE_KEY = iota
	RSA_PUBLIC_KEY
)

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

func EncryptPEM(p *pem.Block, pwd string) (*pem.Block, error) {
	return x509.EncryptPEMBlock(
		rand.Reader,
		p.Type,
		p.Bytes,
		[]byte(pwd),
		x509.PEMCipherAES256,
	)
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

func SaveKey(f string, p *pem.Block, perm os.FileMode) error {
	return ioutil.WriteFile(f, pem.EncodeToMemory(p), perm)
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
