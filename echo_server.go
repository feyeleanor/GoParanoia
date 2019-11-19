package main

import "encoding/pem"
import "fmt"
import "net/http"

func GetPublicKey(w http.ResponseWriter, p *pem.Block, f func()) {
  f()
  fmt.Fprint(w,
    EncodeToString(
      pem.EncodeToMemory(p)))
}
