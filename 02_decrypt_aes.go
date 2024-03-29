package main

import "crypto/aes"
import "crypto/cipher"
import "fmt"
import "os"
import "strconv"

func main() {
  k := os.Getenv("AES_KEY")
  b := read_bytes(os.Args[1:])
  if m, e := Decrypt(k, b); e == nil {
    fmt.Println(string(m))
  } else {
    fmt.Printf("error: %v\n", m)
  }
}

func read_bytes(s []string) (r []byte) {
  for _, v := range s {
    i, _ := strconv.Atoi(v)
    r = append(r, byte(i))
  }
  return
}

func Unpack(m []byte) (iv, r []byte) {
  return m[:aes.BlockSize], m[aes.BlockSize:]
}

func Decrypt(k string, m []byte) (r []byte, e error) {
  var b cipher.Block
  if b, e = aes.NewCipher([]byte(k)); e == nil {
    var iv []byte
    iv, m = Unpack(m)
    r = make([]byte, len(m))
    cipher.
      NewCBCDecrypter(b, iv).
        CryptBlocks(r, m)
  }
  return
}
