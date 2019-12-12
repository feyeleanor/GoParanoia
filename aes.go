package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "fmt"
import "strings"

type AES_channel struct{ ko, ki string }

func (a *AES_channel) EncryptMessage(m string) string {
	b, e := AES_Encrypt(a.ko, m)
	if e != nil {
		fmt.Println("EncryptMessage:", e)
		fmt.Println("EncryptMessage:", []byte(a.ko))
	}
	return EncodeToString(b)
}

func (a *AES_channel) DecryptMessage(m string) (r string) {
	r = read_base64(m)
	var e error
	r, e = AES_Decrypt(a.ki, r)
	if e != nil {
		fmt.Println("DecryptMessage:", e)
	}
	return
}

func (a *AES_channel) EncryptKey(m string) string {
	b, e := AES_Encrypt(a.ko, m)
	if e != nil {
		fmt.Println("EncryptKey:", e)
		fmt.Println("EncryptKey:", []byte(a.ko))
	}
	return EncodeToString(b)
}

func AES_MakeKey(n int) string {
	s := make([]byte, n)
	_, e := rand.Read(s)
	ExitOnError(e, NOT_ENOUGH_RANDOMNESS)
	return string(s)
}

func AES_Encrypt(k, m string) (o []byte, e error) {
	if o, e = PaddedBuffer([]byte(m)); e == nil {
		var b cipher.Block

		if b, e = aes.NewCipher([]byte(k)); e == nil {
			o, e = CryptBlocks(o, b)
		}
	}
	return
}

func AES_Decrypt(k, s string) (r string, e error) {
	var b cipher.Block
	if b, e = aes.NewCipher([]byte(k)); e == nil {
		iv, m := Unpack(s)
		x := make([]byte, len(m))
		cipher.
			NewCBCDecrypter(b, iv).
			CryptBlocks(x, m)
		r = string(x)
	}
	return
}

func AES_DecryptAndTrim(k, s string) (r string, e error) {
	if r, e = AES_Decrypt(k, s); e == nil {
		r = strings.TrimRight(r, "\000")
	}
	return
}

func PaddedBuffer(m []byte) (b []byte, e error) {
	p := len(m) % aes.BlockSize
	b = make([]byte, len(m)+aes.BlockSize-p)
	copy(b, m)
	return
}

func CryptBlocks(b []byte, c cipher.Block) (o []byte, e error) {
	o = make([]byte, aes.BlockSize+len(b))
	var iv []byte
	if iv, e = IV(); e == nil {
		copy(o, iv)
		cipher.
			NewCBCEncrypter(c, o[:aes.BlockSize]).
			CryptBlocks(o[aes.BlockSize:], b)
	}
	return
}

func IV() (b []byte, e error) {
	b = make([]byte, aes.BlockSize)
	_, e = rand.Read(b)
	return
}

func Unpack(s string) (iv, r []byte) {
	r = []byte(s)
	return r[:aes.BlockSize], r[aes.BlockSize:]
}
