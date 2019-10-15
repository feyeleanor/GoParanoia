package main

import "crypto/aes"
import "crypto/cipher"
import "encoding/base64"
import "fmt"
import "os"

func main() {
  k := os.Getenv("AES_KEY")
  s := read_base64(os.Args[1])
  if m, e := Decrypt(k, s); e == nil {
    fmt.Println(m)
  } else {
    fmt.Printf("error: %v\n", m)
  }
}

func read_base64(s string) string {
  b, _ := base64.StdEncoding.DecodeString(s)
  return string(b)
}

func Unpack(s string) (iv, r []byte) {
  m := []byte(s)
  return m[:aes.BlockSize], m[aes.BlockSize:]
}

func Decrypt(k, s string) (r string, e error) {
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
