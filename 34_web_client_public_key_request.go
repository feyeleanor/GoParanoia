package main

import "crypto/rsa"
import "fmt"
import "net/http"
import "os"

type Person string

func (p Person) Report(m ...interface{}) {
  i := []interface{} { p }
  fmt.Println(append(i, m...)...)
}

const KEY = "key"

func main() {
  Person("Alice").Report("received public key:",
    RequestPublicKey(
      ServerAddress(HTTP_ADDRESS), os.Args[1]))
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
  var k interface{}
  var s string

  r, e := http.Get(HTTP_URL(a, KEY, n))
  ExitOnError(e, WEB_REQUEST_FAILED)

  if s = HTTP_readbody(r.Body); s == "" {
    os.Exit(WEB_NO_BODY)
  }

  k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, s, "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
  return k.(*rsa.PublicKey)
}
