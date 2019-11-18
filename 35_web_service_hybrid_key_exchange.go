package main

import "bytes"
import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "io"
import "net/http"
import "os"
import "time"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const KEY = "/key/"

func init() {
  k, e := PEM_Load(RSA_PRIVATE_KEY, os.Args[1], "")
	ExitOnError(e, INVALID_PRIVATE_KEY)

  priv := k.(*rsa.PrivateKey)
  p := PEM_Create(priv.PublicKey)
	http.HandleFunc(KEY, func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      fmt.Println("Bob received request for public key from", r.RemoteAddr)
      fmt.Fprint(w,
        EncodeToString(
          pem.EncodeToMemory(p)))

    case http.MethodPost:
      var s string

      n := r.URL.Path[len(KEY):]
      if len(n) == 0 {
        http.Error(w, "missing nonce", 500)
        return
      }

      s = HTTP_readbody(r.Body)
      if s, e = OAEP_Decrypt(priv, read_base64(s), n); e != nil {
        http.Error(w, "decryption failed", 500)
        return
      }
      fmt.Println("Bob received symmetric key:", s)
    }
  })
}

func main() {
  a := os.Getenv("HTTP_ADDRESS")
	if a == "" {
    a = DEFAULT_ADDRESS
  }

  go func() {
    http.ListenAndServe(a, nil)
  }()
  time.Sleep(2 * time.Second)

  n := os.Args[2]
  k := os.Args[3]
  p := RequestPublicKey(a, n)
  fmt.Println("Alice received public key:", p)
  fmt.Println("Alice sends symmetric key:", k)
  SendSymmetricKey(p, a, k, n)
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
  var k interface{}
  var s string

  r, e := http.Get(HTTP + a + KEY + n)
  ExitOnError(e, WEB_REQUEST_FAILED)

  if s = HTTP_readbody(r.Body); s == "" {
    os.Exit(WEB_NO_BODY)
  }

  k, e = PEM_ReadBase64(RSA_PUBLIC_KEY, s, "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
  return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, a, k, n string) {
	b, e := OAEP_Encrypt(pk, k, n)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)

  _, e = http.Post(
    HTTP + a + KEY + n,
    "application/octet-stream",
    EncodeToIOReader(b))

  ExitOnError(e, WEB_REQUEST_FAILED)
}

func EncodeToIOReader(m []byte) io.Reader {
	return bytes.NewBufferString(EncodeToString(m))
}
