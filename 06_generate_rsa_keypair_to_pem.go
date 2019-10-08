package main

import "crypto/rand"
import "crypto/rsa"
import "crypto/x509"
import "encoding/pem"
import "io/ioutil"
import "os"
import "strconv"

const (
  _ = iota
  MISSING_FILENAME
  INVALID_KEYSIZE
  CREATE_KEY_FAILED
  PEM_ENCRYPTION_FAILED
  FILE_WRITE_FAILED
)
const DEFAULT_KEYSIZE = 1024

func main() {
  var f string
  var n uint64
  var e error
  var k *rsa.PrivateKey

  if f = os.Args[1]; f == "" {
    os.Exit(MISSING_FILENAME)
  }
  if n, e = strconv.ParseUint(os.Args[2], 10, 64); e != nil {
    os.Exit(INVALID_KEYSIZE)
  }
  if n == 0 {
    n = DEFAULT_KEYSIZE
  }
  if k, e = CreatePrivateKey(int(n)); e != nil {
    os.Exit(CREATE_KEY_FAILED)
  }
  p := CreatePEM(k)
  if pwd := os.Getenv("PEM_KEY"); pwd != "" {
    if p, e = EncryptPEM(p, pwd); e != nil {
      os.Exit(PEM_ENCRYPTION_FAILED)
    }
  }
  if e = SaveKey(f, p, 0644); e != nil {
    os.Exit(FILE_WRITE_FAILED)
  }
}

func SaveKey(f string, p *pem.Block, perm os.FileMode) error {
  return ioutil.WriteFile(f, pem.EncodeToMemory(p), perm)
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
