package main

import "bytes"
import "crypto/rsa"
import "io"
import "net/http"
import "os"

const KEY = "key"
const ALICE Person = "Alice"

func main() {
  a := ServerAddress(HTTP_ADDRESS)
  n := os.Args[1]
  k := os.Args[2]
  p := RequestPublicKey(a, n)
  ALICE.Report("received public key:", p)
  ALICE.Report("sends symmetric key:", k)
  SendSymmetricKey(p, a, k, n)
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

func SendSymmetricKey(pk *rsa.PublicKey, a, k, n string) {
	b, e := OAEP_Encrypt(pk, k, n)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)

  _, e = http.Post(
    HTTP_URL(a, KEY, n),
    OCTET_STREAM,
    EncodeToIOReader(b))

  ExitOnError(e, WEB_REQUEST_FAILED)
}

func EncodeToIOReader(m []byte) io.Reader {
	return bytes.NewBufferString(EncodeToString(m))
}
