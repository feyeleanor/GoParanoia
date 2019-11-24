package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"

type Person string

func (p Person) Report(m ...interface{}) {
  i := []interface{} { p }
  fmt.Println(append(i, m...)...)
}

const KEY = "key"

func init() {
  k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

  p := PEM_Create(k.(*rsa.PrivateKey).PublicKey)
	HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
    Person("Bob").Report("received nonce:", SubPath(KEY, r))
    fmt.Fprint(w,
      EncodeToString(
        pem.EncodeToMemory(p)))
  })
}

func main() {
  http.ListenAndServe(
    ServerAddress(HTTP_ADDRESS), nil)
}
