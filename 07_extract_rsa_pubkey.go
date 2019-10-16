package main

import "crypto/rsa"
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
  NOT_A_PRIVATE_KEY
  PEM_PASSWORD_REQUIRED
  PEM_DECRYPTION_FAILED
  INVALID_PRIVATE_KEY
  FILE_WRITE_FAILED
)

func main() {
  var e error
  var k *rsa.PrivateKey

  f := os.Args[1]
  b := LoadPEM(f)
  k, e = x509.ParsePKCS1PrivateKey(b)
  ExitOnError(e, INVALID_PRIVATE_KEY)

  f = f + ".pub"
  e = SaveKey(f, CreatePEM(k.PublicKey), 0644)
  ExitOnError(e, FILE_WRITE_FAILED)
}

func LoadPEM(s string) []byte {
  return DecryptPEM(
    DecodePEM(
      LoadFile(s)))
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

func SaveKey(f string, p *pem.Block, perm os.FileMode) error {
  return ioutil.WriteFile(f, pem.EncodeToMemory(p), perm)
}

func CreatePEM(k rsa.PublicKey) *pem.Block {
  return &pem.Block{
    Type:  "RSA PUBLIC KEY",
    Bytes: x509.MarshalPKCS1PublicKey(&k),
  }
}

func DecodePEM(b []byte) (p *pem.Block) {
  switch p, _ = pem.Decode(b); {
  case p == nil:
    os.Exit(NOT_A_PEM_FILE)
  case p.Type != "RSA PRIVATE KEY":
    os.Exit(NOT_A_PRIVATE_KEY)
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
