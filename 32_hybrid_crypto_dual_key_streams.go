package main

import "crypto/rand"
import "crypto/rsa"
import "encoding/pem"
import "fmt"
import "os"

func main() {
	AtoB := make(chan string)
  BtoA := Launch(AtoB, func(in, out chan string) {
    kp, p := LoadKeys(os.Args[1], "")
    ki, ko := ServerHandshake(kp, p, in, out)

    Transmitter(ki, os.Args[4:], func(k, v string) {
  	  fmt.Println("Bob wants to say:", v)
      out <- EncryptMessage(ko, v)
			fmt.Println("Bob heard:", DecryptMessage(k, <-in))
		})
		close(out)
  })

  n := os.Args[2]
  ki := os.Args[3]
  ko := ClientHandshake(ki, n, BtoA, AtoB)
  fmt.Println(
    "Alice received symmetric key:",
    EncodeToString([]byte(ko)))

  Receiver(ki, BtoA, func(v string) {
		fmt.Println("Alice heard:", v)
		v = fmt.Sprintf("%v received", v)
	  fmt.Println("Alice wants to say:", v)
    AtoB <- EncryptMessage(ko, v)
  })
}

func ClientHandshake(ki, n string, in, out chan string) string {
	SendSymmetricKey(
    RequestPublicKey(in, out, n), out, ki, n)

  return DecryptMessage(ki, <- in)
}

func ServerHandshake(kp *rsa.PrivateKey, p *pem.Block, in, out chan string) (ki, ko string) {
	n := <-in
	fmt.Println("Server received nonce:", n)

	out <- EncodeToString(pem.EncodeToMemory(p))
  ko = ReceiveSymmetricKey(kp, in, n)
	fmt.Println("Bob received symmetric key:", ko)

  b := make([]byte, 32)
  _, e := rand.Read(b)
  ExitOnError(e, NOT_ENOUGH_RANDOMNESS)
  ki = string(b)
	out <- EncryptMessage(ko, ki)
  return
}

func Launch(in chan string, f func(chan string, chan string)) (out chan string) {
  out = make(chan string)
  go func(i, o chan string) {
    f(i, o)
  }(in, out)
  return
}

func LoadKeys(f, pwd string) (*rsa.PrivateKey, *pem.Block) {
  k, e := PEM_Load(RSA_PRIVATE_KEY, f, pwd)
	ExitOnError(e, INVALID_PRIVATE_KEY)

  kp := k.(*rsa.PrivateKey)
	p := PEM_Create(kp.PublicKey)
  return kp, p
}

func RequestPublicKey(i, o chan string, n string) *rsa.PublicKey {
	o <- n
  k, e := PEM_ReadBase64(RSA_PUBLIC_KEY, <-i, "")
	ExitOnError(e, INVALID_PUBLIC_KEY)
  return k.(*rsa.PublicKey)
}

func SendSymmetricKey(pk *rsa.PublicKey, out chan string, k, l string) {
	k = EncodeToString([]byte(k))
	b, e := OAEP_Encrypt(pk, k, l)
	ExitOnError(e, RSA_ENCRYPTION_FAILED)
	out <- string(b)
}

func ReceiveSymmetricKey(pk *rsa.PrivateKey, in chan string, n string) string {
	k, e := OAEP_Decrypt(pk, <-in, n)
	ExitOnError(e, RSA_DECRYPTION_FAILED)
  return read_base64(k)
}

func Transmitter(k string, m []string, f func(k, v string)) {
  for _, v := range m {
    f(k, v)
  }
}

func Receiver(k string, in chan string, f func(string)) {
  for v := range in {
    f(DecryptMessage(k, v))
  }
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
