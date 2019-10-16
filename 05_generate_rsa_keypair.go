package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/base64"
import "encoding/pem"
import "fmt"
import "os"
import "strconv"

const (
  _ = iota
  MISSING_KEYSIZE
  INVALID_KEYSIZE
  CREATE_KEY_FAILED
  PEM_ENCRYPTION_FAILED
)

func main() {
  var n int
  var e error
  var k *rsa.PrivateKey

  if len(os.Args) < 2 {
    os.Exit(MISSING_KEYSIZE)
  }
  if n, e = strconv.Atoi(os.Args[1]); n == 0 {
    os.Exit(INVALID_KEYSIZE)
  }
  k, e = CreatePrivateKey(n)
  ExitOnError(e, CREATE_KEY_FAILED)

  p := CreatePEM(k)
  if pwd := os.Getenv("PEM_KEY"); pwd != "" {
    p, e = EncryptPEM(p, pwd)
    ExitOnError(e, PEM_ENCRYPTION_FAILED)
  }
  fmt.Println(PrintableKey(p))
}

func PrintableKey(p *pem.Block) string {
  b := pem.EncodeToMemory(p)
  return base64.StdEncoding.EncodeToString(b)
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

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
