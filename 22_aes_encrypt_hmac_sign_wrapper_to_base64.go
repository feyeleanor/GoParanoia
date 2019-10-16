package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/hmac"
import "crypto/rand"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

const (
  _ = iota
  ENCRYPTION_FAILED
)

func main() {
  var e error
  var m []byte

  k := os.Getenv("AES_KEY")
  s := os.Args[1]

  m, e = Encrypt(s, k)
  ExitOnError(e, ENCRYPTION_FAILED)

  hk := os.Getenv("HMAC_KEY")
  h := hmac.New(sha512.New, []byte(hk))
  h.Write(m)

  fmt.Println(
    EncodeToString(h.Sum(nil)) +
    EncodeToString(m))
}

func EncodeToString(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
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

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
