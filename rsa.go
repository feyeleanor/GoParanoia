package main

import "crypto"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"

func PKCS1v15_Encrypt(k *rsa.PublicKey, m string) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, k, []byte(m))
}

func PKCS1v15_Decrypt(k *rsa.PrivateKey, m string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptPKCS1v15(rand.Reader, k, []byte(m))
	r = string(b)
	return
}

func PKCS1v15_Sign(k *rsa.PrivateKey, s string) ([]byte, error) {
	h := sha512.Sum512([]byte(s))
	return rsa.SignPKCS1v15(
		rand.Reader,
		k,
		crypto.SHA512,
		h[:])
}

func PKCS1v15_Verify(k *rsa.PublicKey, s, m string) error {
	h := sha512.Sum512([]byte(m))
	return rsa.VerifyPKCS1v15(
		k,
		crypto.SHA512,
		h[:],
		[]byte(s))
}

func OAEP_Encrypt(k *rsa.PublicKey, m, l string) ([]byte, error) {
	return rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
}

func OAEP_Decrypt(k *rsa.PrivateKey, m, l string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
	r = string(b)
	return
}

func PSS_Sign(k *rsa.PrivateKey, m string, o *rsa.PSSOptions) (r []byte, e error) {
	hs := sha512.Sum512([]byte(m))
	return rsa.SignPSS(rand.Reader, k, crypto.SHA512, hs[:], o)
}

func PSS_Verify(k *rsa.PublicKey, s, m string) error {
	h := sha512.Sum512([]byte(m))
	return rsa.VerifyPSS(k, crypto.SHA512, h[:], []byte(s), nil)
}
