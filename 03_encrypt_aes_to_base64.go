package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "encoding/base64"
import "fmt"
import "os"

func main() {
  k := os.Getenv("AES_KEY")
  if m, e := Encrypt(k, os.Args[1]); e == nil {
    PrintEncrypted(m)
  } else {
    fmt.Printf("error: %v\n", e)
  }
}

func PrintEncrypted(m []byte) {
  fmt.Println(base64.StdEncoding.EncodeToString(m))
}

func Encrypt(k, m string) (o []byte, e error) {
  if o, e = PaddedBuffer([]byte(m)); e == nil {
    var b cipher.Block

    if b, e = aes.NewCipher([]byte(k)); e == nil {
      o, e = CryptBlocks(o, b)
    }
  }
  return
}

func PaddedBuffer(m []byte) (b []byte, e error) {
  p := len(m) % aes.BlockSize
  b = make([]byte, len(m) + aes.BlockSize - p)
  copy(b, m)
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
