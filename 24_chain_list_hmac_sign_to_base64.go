package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

func main() {
  var s *SignedList

  k := os.Getenv("HMAC_KEY")
  for i := len(os.Args) - 1; i > 0; i-- {
    s = s.Push(k, os.Args[i])
  }
  s.Each(
    func(s SignedList) {
      fmt.Println(s.H, s.V) })
}

type SignedList struct {
  V string
  H string
  *SignedList
}

func (s *SignedList) Push(k, v string) *SignedList {
  h := hmac.New(sha512.New, []byte(k))
  if s != nil {
    h.Write([]byte(s.H))
  }
  h.Write([]byte(v))
  return &SignedList{
    v, EncodeToString(h.Sum(nil)), s,
  }
}

func (s *SignedList) Each(f func(SignedList)) {
  if s != nil {
    f(*s)
    s.SignedList.Each(f)
  }
  return
}

func EncodeToString(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}
