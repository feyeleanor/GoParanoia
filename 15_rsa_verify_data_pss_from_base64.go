package main

import "crypto"
import "crypto/rsa"
import "crypto/sha512"
import "encoding/base64"
import "crypto/x509"
import "encoding/pem"
import "fmt"
import "io/ioutil"
import "os"

const (
  _ = iota
  MISSING_FILENAME
  FILE_UNREADABLE
  NOT_A_PEM_FILE
  NOT_A_PUBLIC_KEY
  PEM_PASSWORD_REQUIRED
  PEM_DECRYPTION_FAILED
  INVALID_PUBLIC_KEY
  VERIFICATION_FAILED
)

func main() {
  var e error
  var k *rsa.PublicKey

  f := os.Args[1]
  m := []byte(os.Args[2])

  hs := sha512.Sum512(m)
  s := read_base64(os.Args[3])
  b := LoadFile(f)
  p := DecodePEM(b)
  b = DecryptPEM(p)

  k, e = x509.ParsePKCS1PublicKey(b)
  ExitOnError(e, INVALID_PUBLIC_KEY)

  if e = rsa.VerifyPSS(k, crypto.SHA512, hs[:], s, nil); e != nil {
    fmt.Println("Signature Verification Failed")
    os.Exit(VERIFICATION_FAILED)
  }
  fmt.Println("Signature Verification Succeeded")
}

func read_base64(s string) []byte {
  b, _ := base64.StdEncoding.DecodeString(s)
  return b
}

func LoadFile(s string) (b []byte) {
  var e error
  if s == "" {
    os.Exit(MISSING_FILENAME)
  }

  b, e = ioutil.ReadFile(s)
  ExitOnError(e, FILE_UNREADABLE)
  return
}

func DecodePEM(b []byte) (p *pem.Block) {
  switch p, _ = pem.Decode(b); {
  case p == nil:
    os.Exit(NOT_A_PEM_FILE)
  case p.Type != "RSA PUBLIC KEY":
    os.Exit(NOT_A_PUBLIC_KEY)
  }
  return
}

func DecryptPEM(p *pem.Block) (b []byte) {
  if x509.IsEncryptedPEMBlock(p) {
	if pwd := os.Getenv("PEM_KEY"); pwd != "" {
      var e error
      b, e = x509.DecryptPEMBlock(p, []byte(pwd))
      ExitOnError(e, PEM_DECRYPTION_FAILED)
    } else {
      os.Exit(PEM_PASSWORD_REQUIRED)
    }
  } else {
    b = p.Bytes
  }
  return
}

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
