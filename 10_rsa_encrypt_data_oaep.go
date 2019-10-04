
package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "crypto/x509"
import "encoding/pem"
import "fmt"
import "io/ioutil"
import "os"

const (
  MISSING_FILENAME = 1
  FILE_UNREADABLE = iota
  NOT_A_PEM_FILE = iota
  NOT_A_PUBLIC_KEY = iota
  PEM_PASSWORD_REQUIRED = iota
  PEM_DECRYPTION_FAILED = iota
  INVALID_PUBLIC_KEY = iota
  ENCRYPTION_FAILED = iota
)

func main() {
  var e error
  var k *rsa.PublicKey

  f := os.Args[1]
  l := []byte(os.Args[2])
  m := []byte(os.Args[3])
  b := LoadFile(f)
  p := DecodePEM(b)
  b = DecryptPEM(p)

  if k, e = x509.ParsePKCS1PublicKey(b); e != nil {
    os.Exit(INVALID_PUBLIC_KEY)
  }

  if b, e = rsa.EncryptOAEP(sha512.New(), rand.Reader, k, m, l); e != nil {
    fmt.Println(e)
    os.Exit(ENCRYPTION_FAILED)
  }
  fmt.Println(b)
}

func LoadFile(s string) (b []byte) {
  var e error
  if s == "" {
    os.Exit(MISSING_FILENAME)
  }

  if b, e = ioutil.ReadFile(s); e != nil {
    os.Exit(FILE_UNREADABLE)
  }
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
      if b, e = x509.DecryptPEMBlock(p, []byte(pwd)); e != nil {
        os.Exit(PEM_DECRYPTION_FAILED)
      }
    } else {
      os.Exit(PEM_PASSWORD_REQUIRED)
    }
  } else {
    b = p.Bytes
  }
  return
}
