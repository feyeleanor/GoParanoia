package main

import (
  "fmt"
  "crypto/sha512"
  "encoding/base64"
  "os"
)

func main() {
  m := os.Args[1]
  h := sha512.New()
  h.Write([]byte(m))
  s := h.Sum(nil)
  fmt.Printf("SHA512 Hex    : %x\n", s)
  fmt.Printf("SHA512 Base64 : %v\n", base64.StdEncoding.EncodeToString(s))
}
