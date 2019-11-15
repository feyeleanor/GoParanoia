package main

import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "net/http"
import "os"
import "strings"
import "time"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const KEY = "/key/"
const MESSAGE = "/message/"

var sessions map[string] string

func init() {
  sessions = make(map[string] string)
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

      if s, e = HTTP_readbody(r.Body); e != nil {
        http.Error(w, "missing symmetric key", 500)
        return
      }

      if s, e = OAEP_Decrypt(priv, read_base64(s), n); e != nil {
        http.Error(w, "decryption failed", 500)
        return
      }
      fmt.Println("Bob received symmetric key:", s)
      sessions[n] = s
    }
  })

  http.HandleFunc(MESSAGE, func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPut:
      n := r.URL.Path[len(MESSAGE):]
      s := sessions[n]
      if len(s) == 0 {
        http.Error(w, "unknown nonce", 500)
      }

      m, e := HTTP_readbody(r.Body)
      if e != nil {
        http.Error(w, "missing message", 500)
        return
      }

      m = DecryptMessage(s, m)
      fmt.Println("Bod heard:", m)

      m = fmt.Sprintf("%v received", m)
      fmt.Println("Bob wants to say:", m)
      fmt.Fprint(w, EncryptMessage(s, m))
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

  for _, m := range os.Args[4:] {
    fmt.Println("Alice wants to say:", m)
    fmt.Println("Alice heard:", SendMessage(a, n, k, m))
  }
}

func RequestPublicKey(a string, n string) *rsa.PublicKey {
  var k interface{}
  var s string

  r, e := http.Get(HTTP + a + KEY + n)
  ExitOnError(e, WEB_REQUEST_FAILED)

  s, e = HTTP_readbody(r.Body)
	ExitOnError(e, WEB_NO_BODY)

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
    EncodeToReader(b))

  ExitOnError(e, WEB_REQUEST_FAILED)
}

func SendMessage(a, n, k, m string) string {
  r, e := Put(HTTP + a + MESSAGE + n, EncryptMessage(k, m))
  ExitOnError(e, WEB_REQUEST_FAILED)

  m, e = HTTP_readbody(r.Body)
	ExitOnError(e, WEB_NO_BODY)

  return DecryptMessage(k, m)
}

func DecryptMessage(k, v string) string {
	v = read_base64(v)
	r, e := AESDecrypt(k, v)
	ExitOnError(e, AES_DECRYPTION_FAILED)
  return r
}

func EncryptMessage(k, v string) string {
	b, e := AESEncrypt(k, v)
	ExitOnError(e, AES_ENCRYPTION_FAILED)
  return EncodeToString(b)
}

func Put(url, m string) (*http.Response, error) {
	r, e := http.NewRequest("PUT", url, strings.NewReader(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	r.ContentLength = int64(len(m))
	return http.DefaultClient.Do(r)
}
