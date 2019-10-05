package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/pem"
import "fmt"
import "os"
import "strconv"

const (
  _ = iota
  MISSING_KEYSIZE
  CREATE_KEY_FAILED
  PEM_ENCRYPTION_FAILED
)

func main() {
  var n int
  var e error
  var k *rsa.PrivateKey

  if n, e = strconv.Atoi(os.Args[1]); e != nil {
    os.Exit(MISSING_KEYSIZE)
  }
  if k, e = CreatePrivateKey(n); e != nil {
    os.Exit(CREATE_KEY_FAILED)
  }
  p := CreatePEM(k)
  if pwd := os.Getenv("PEM_KEY"); pwd != "" {
    if p, e = EncryptPEM(p, pwd); e != nil {
      os.Exit(PEM_ENCRYPTION_FAILED)
    }
  }
  fmt.Println(pem.EncodeToMemory(p))
}

func CreatePrivateKey(n int) (*rsa.PrivateKey, error) {
  return rsa.GenerateKey(rand.Reader, n)
}

func CreatePEM(k *rsa.PrivateKey) *pem.Block {
  return &pem.Block{
    Type:  "RSA PRIVATE KEY",
    Bytes: x509.MarshalPKCS1PrivateKey(k),
  }
}

func EncryptPEM(p *pem.Block, s string) (*pem.Block, error) {
  return x509.EncryptPEMBlock(
    rand.Reader,
    p.Type,
    p.Bytes,
    []byte(s),
    x509.PEMCipherAES256,
  )
}
