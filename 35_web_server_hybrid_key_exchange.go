package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"

const KEY = "key"
const BOB Person = "Bob"

func init() {
  k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

  priv := k.(*rsa.PrivateKey)
  p := PEM_Create(priv.PublicKey)
	HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      BOB.Report("received request for public key from", r.RemoteAddr)
      fmt.Fprint(w,
        EncodeToString(
          pem.EncodeToMemory(p)))

    case http.MethodPost:
      if n := SubPath(KEY, r); len(n) == 0 {
        http.Error(w, "missing nonce", 500)
      } else {
        s := HTTP_readbody(r.Body)
        if s, e = OAEP_Decrypt(priv, read_base64(s), n); e != nil {
          http.Error(w, "decryption failed", 500)
        } else {
          BOB.Report("received symmetric key:", s)
        }
      }
    }
  })
}

func main() {
  http.ListenAndServe(
    ServerAddress(HTTP_ADDRESS), nil)
}
