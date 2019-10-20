package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "crypto/x509"
import "encoding/pem"
import "fmt"
import "io/ioutil"
import "os"

const (
	_ = iota
	MISSING_FILENAME
	FILE_UNREADABLE
	NOT_A_PEM_FILE
	NOT_AN_RSA_KEY
	PEM_PASSWORD_REQUIRED
	PEM_DECRYPTION_FAILED
	INVALID_PUBLIC_KEY
	INVALID_PRIVATE_KEY
	ENCRYPTION_FAILED
	DECRYPTION_FAILED
)

func main() {
	AtoB := make(chan string)
	BtoA := make(chan string)

	go func() {
		var e error
		var kBob *rsa.PrivateKey

		fBob := LoadFile(os.Args[1])
		b := DecodePEM(fBob).Bytes

		kBob, e = x509.ParsePKCS1PrivateKey(b)
		ExitOnError(e, INVALID_PRIVATE_KEY)

		p := CreatePEM(kBob.PublicKey)
		l := <-AtoB
		fmt.Println("Bob heard:", l)

		BtoA <- EncodeToString(pem.EncodeToMemory(p))
		m := <-AtoB

		var k string
		k, e = RSADecrypt(kBob, m, l)
		ExitOnError(e, DECRYPTION_FAILED)

		fmt.Println("Bob heard:", k)
		k = read_base64(k)

		for _, v := range os.Args[4:] {
			SendMessageFrom("Bob", BtoA, k, v)
			v = read_base64(<-AtoB)
			v, e = AESDecrypt(k, v)
			ExitOnError(e, DECRYPTION_FAILED)
			fmt.Println("Bob heard:", v)
		}
		close(BtoA)
	}()

	l := os.Args[2]
	k := os.Args[3]
	kBob := RequestPublicKey(AtoB, BtoA, l)
	SendSymmetricKey(AtoB, kBob, k, l)

	var e error
	for v := range BtoA {
		v = read_base64(v)
		v, e = AESDecrypt(k, v)
		ExitOnError(e, DECRYPTION_FAILED)
		fmt.Println("Alice heard:", v)
		v = fmt.Sprintf("%v received", v)
		SendMessageFrom("Alice", AtoB, k, v)
	}
}

func RequestPublicKey(out, in chan string, m string) (k *rsa.PublicKey) {
	var e error

	out <- m
	m = read_base64(<-in)
	b := DecodePEM([]byte(m)).Bytes
	k, e = x509.ParsePKCS1PublicKey(b)
	ExitOnError(e, INVALID_PUBLIC_KEY)
	return
}

func SendSymmetricKey(out chan string, pk *rsa.PublicKey, k, l string) {
	k = EncodeToString([]byte(k))
	b, e := RSAEncrypt(pk, k, l)
	ExitOnError(e, ENCRYPTION_FAILED)
	out <- string(b)
}

func SendMessageFrom(n string, c chan string, k, v string) {
	fmt.Println(n, "wants to say:", v)
	b, e := AESEncrypt(k, v)
	ExitOnError(e, DECRYPTION_FAILED)
	c <- EncodeToString(b)
}

func RSAEncrypt(k *rsa.PublicKey, m, l string) (b []byte, e error) {
	return rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
}

func RSADecrypt(k *rsa.PrivateKey, m, l string) (r string, e error) {
	var b []byte
	b, e = rsa.DecryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
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
	case p.Type != "RSA PUBLIC KEY" && p.Type != "RSA PRIVATE KEY":
		os.Exit(NOT_AN_RSA_KEY)
	}
	return
}

func CreatePEM(k rsa.PublicKey) *pem.Block {
	return &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&k),
	}
}
