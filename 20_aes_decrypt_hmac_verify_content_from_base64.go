package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

const (
  _ = iota
  DECRYPTION_FAILED
  INCORRECT_CONTENT
  INVALID_SIGNATURE
)

func main() {
  var e error

  hk := os.Getenv("HMAC_KEY")
  h := hmac.New(sha512.New, []byte(hk))

  m := os.Args[1]
  h.Write([]byte(m))

  s := os.Args[2]
  hs := []byte(read_base64(s[0:88]))
  ms := []byte(read_base64(s[88:]))

  k := os.Getenv("AES_KEY")
  switch ms, e = Decrypt(ms, k); {
  case e != nil:
    fmt.Printf("error: %v\n", m)
    os.Exit(DECRYPTION_FAILED)
  case string(ms) != m:
    fmt.Println("error: content doesn't match")
    os.Exit(INCORRECT_CONTENT)
  case !hmac.Equal(h.Sum(nil), hs):
    fmt.Println("Signature Verification Failed")
    os.Exit(INVALID_SIGNATURE)
  default:
    fmt.Println("Signature Verification Succeeded")
  }
}

func read_base64(s string) string {
  b, _ := base64.StdEncoding.DecodeString(s)
  return string(b)
}

func Unpack(m []byte) (iv, r []byte) {
  return m[:aes.BlockSize], m[aes.BlockSize:]
}

func Decrypt(m []byte, k string) (r []byte, e error) {
  var b cipher.Block
  if b, e = aes.NewCipher([]byte(k)); e == nil {
    var iv []byte
    iv, m = Unpack(m)
    r = make([]byte, len(m))
    cipher.
      NewCBCDecrypter(b, iv).
      CryptBlocks(r, m)
    r = TrimBuffer(r)
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
