package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"
import "time"

type Person string

func (p Person) Report(m ...interface{}) {
  i := []interface{} { p }
  fmt.Println(append(i, m...)...)
}

const BOB Person = "Bob"
const ALICE Person = "Alice"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const KEY = "/key/"

const HTTP_ADDRESS = "HTTP_ADDRESS"

func init() {
  k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

  p := PEM_Create(k.(*rsa.PrivateKey).PublicKey)
	http.HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Path[len(KEY):]
    BOB.Report("received nonce:", n)
    fmt.Fprint(w,
      EncodeToString(
        pem.EncodeToMemory(p)))
  })
}

func main() {
  a := os.Getenv(HTTP_ADDRESS)
	if a == "" {
    a = DEFAULT_ADDRESS
  }

  go func() {
    http.ListenAndServe(a, nil)
  }()
  time.Sleep(2 * time.Second)

  ALICE.Report("received public key:", RequestPublicKey(a, os.Args[2]))
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
  var k interface{}
  var s string

  url := HTTP + a + KEY + "%v"
  r, e := http.Get(fmt.Sprintf(url, n))
  ExitOnError(e, WEB_REQUEST_FAILED)

  if s = HTTP_readbody(r.Body); s == "" {
    os.Exit(WEB_NO_BODY)
  }

  k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, s, "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
  return k.(*rsa.PublicKey)
}
