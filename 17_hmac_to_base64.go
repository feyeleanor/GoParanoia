package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

func main() {
  k := os.Getenv("HMAC_KEY")
  h := hmac.New(sha512.New, []byte(k))
  h.Write([]byte(os.Args[1]))
  fmt.Println(EncodeToString(h.Sum(nil)))
}

func EncodeToString(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}
