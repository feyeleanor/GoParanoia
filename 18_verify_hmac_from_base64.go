package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

func main() {
  k := []byte(os.Getenv("HMAC_KEY"))
  m := hmac.New(sha512.New, k)
  m.Write([]byte(os.Args[1]))

  h, _ := base64.StdEncoding.DecodeString(os.Args[2])
  if hmac.Equal(h, m.Sum(nil)) {
    fmt.Println("Signature Verification Succeeded")
  } else {
    fmt.Println("Signature Verification Failed")
    os.Exit(1)
  }
}
