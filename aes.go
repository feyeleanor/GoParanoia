package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"

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
		r = string(TrimBuffer(x))
	}
	return
}

func TrimBuffer(m []byte) (r []byte) {
	r = m
	for i := len(m) - 1; i > 0; i-- {
		if m[i] == 0 {
			r = m[:i]
		}
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
