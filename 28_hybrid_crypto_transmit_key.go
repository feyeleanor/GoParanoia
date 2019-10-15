package main

import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha512"
import "crypto/x509"
import "encoding/base64"
import "encoding/pem"
import "fmt"
import "io/ioutil"
import "os"

const (
  _ = iota
  MISSING_FILENAME
  FILE_UNREADABLE
  NOT_A_PEM_FILE
  NOT_AN_RSA_KEY
  PEM_PASSWORD_REQUIRED
  PEM_DECRYPTION_FAILED
  INVALID_PUBLIC_KEY
  INVALID_PRIVATE_KEY
  ENCRYPTION_FAILED
  DECRYPTION_FAILED
)

func main() {
  AtoB := make(chan string)
  BtoA := make(chan string)

  go func() {
    var e error
    var kBob *rsa.PrivateKey

    fBob := LoadFile(os.Args[1])
    b := DecodePEM(fBob).Bytes

    if kBob, e = x509.ParsePKCS1PrivateKey(b); e != nil {
      os.Exit(INVALID_PRIVATE_KEY)
    }
    p := CreatePEM(kBob.PublicKey)

    l := <- AtoB
    fmt.Println("Bob heard:", l)

    BtoA <- EncodeToString(pem.EncodeToMemory(p))
    m := <- AtoB

    var k string
    if k, e = RSADecrypt(kBob, m, l); e != nil {
      fmt.Println(e)
      os.Exit(DECRYPTION_FAILED)
    }
    fmt.Println("Bob heard:", k)
    k = read_base64(k)

    for _, v := range os.Args[4:] {
      fmt.Println("Bob wants to say:", v)
      if b, e = AESEncrypt(k, v); e != nil {
        fmt.Println(e)
        os.Exit(DECRYPTION_FAILED)
      }
      BtoA <- EncodeToString(b)
    }
    close(BtoA)
  }()

  l := os.Args[2]
  k := os.Args[3]
  kBob := RequestPublicKey(AtoB, BtoA, l)
  SendSymmetricKey(AtoB, kBob, k, l)

  var e error
  for v := range BtoA {
    v = read_base64(v)
    if v, e = AESDecrypt(k, v); e != nil {
      os.Exit(DECRYPTION_FAILED)
    }
    fmt.Println("Alice heard:", v)
  }
}

func RequestPublicKey(out, in chan string, m string) (k *rsa.PublicKey) {
  var e error

  out <- m
  m = read_base64(<- in)
  b := DecodePEM([]byte(m)).Bytes
  if k, e = x509.ParsePKCS1PublicKey(b); e != nil {
    os.Exit(INVALID_PUBLIC_KEY)
  }
  return
}

func SendSymmetricKey(out chan string, pk *rsa.PublicKey, k, l string) {
  k = EncodeToString([]byte(k))
  if b, e := RSAEncrypt(pk, k, l); e == nil {
    out <- string(b)
  } else {
    fmt.Println(e)
    os.Exit(ENCRYPTION_FAILED)
  }
}

func read_base64(s string) string {
  b, _ := base64.StdEncoding.DecodeString(s)
  return string(b)
}

func EncodeToString(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}

func RSAEncrypt(k *rsa.PublicKey, m, l string) (b []byte, e error) {
  return rsa.EncryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
}

func RSADecrypt(k *rsa.PrivateKey, m, l string) (r string, e error) {
  var b []byte
  b, e = rsa.DecryptOAEP(sha512.New(), rand.Reader, k, []byte(m), []byte(l))
  r = string(b)
  return
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
  case p.Type != "RSA PUBLIC KEY" && p.Type != "RSA PRIVATE KEY":
    os.Exit(NOT_AN_RSA_KEY)
  }
  return
}

func CreatePEM(k rsa.PublicKey) *pem.Block {
  return &pem.Block{
    Type:  "RSA PUBLIC KEY",
    Bytes: x509.MarshalPKCS1PublicKey(&k),
  }
}

func Unpack(s string) (iv, r []byte) {
  m := []byte(s)
  return m[:aes.BlockSize], m[aes.BlockSize:]
}

func AESDecrypt(k, s string) (r string, e error) {
  var b cipher.Block
  if b, e = aes.NewCipher([]byte(k)); e == nil {
    iv, m := Unpack(s)
    x := make([]byte, len(m))
    cipher.
      NewCBCDecrypter(b, iv).
      CryptBlocks(x, m)
    r = string(x)
  }
  return
}

func AESEncrypt(k, m string) (o []byte, e error) {
  if o, e = PaddedBuffer([]byte(m)); e == nil {
    var b cipher.Block

    if b, e = aes.NewCipher([]byte(k)); e == nil {
      o, e = CryptBlocks(o, b)
    }
  }
  return
}

func PaddedBuffer(m []byte) (b []byte, e error) {
  p := len(m) % aes.BlockSize
  b = make([]byte, len(m) + aes.BlockSize - p)
  copy(b, m)
  return
}

func CryptBlocks(b []byte, c cipher.Block) (o []byte, e error) {
  o = make([]byte, aes.BlockSize + len(b))
  var iv []byte
  if iv, e = IV(); e == nil {
    copy(o, iv)
  	cipher.
  	  NewCBCEncrypter(c, o[:aes.BlockSize]).
  	  CryptBlocks(o[aes.BlockSize:], b)
  }
  return
}

func IV() (b []byte, e error) {
  b = make([]byte, aes.BlockSize)
  _, e = rand.Read(b)
  return
}
