package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "fmt"
import "os"

func main() {
  if m, e := Encrypt(os.Args[1], os.Getenv("AES_KEY")); e == nil {
    fmt.Println(m)
  } else {
  	fmt.Printf("error: %v\n", e)
  }
}

func Encrypt(m, k string) (o []byte, e error) {
  if o, e = PaddedBuffer([]byte(m)); e == nil {
    var b cipher.Block

    if b, e = aes.NewCipher([]byte(k)); e == nil {
      o, e = CryptBlocks(o, b)
    }
  }
  return
}

func PaddedBuffer(m []byte) (b []byte, e error) {
  b = append(b, m...)
  if p := len(b) % aes.BlockSize; p != 0 {
    p = aes.BlockSize - p
    b = append(b, make([]byte, p)...)  // padding with NUL!!!!
  }
  return
}

func CryptBlocks(b []byte, c cipher.Block) (o []byte, e error) {
  o = make([]byte, aes.BlockSize + len(b)) 
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