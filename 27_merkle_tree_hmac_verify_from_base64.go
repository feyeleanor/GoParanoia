package main

import "crypto/hmac"
import "crypto/sha512"
import "encoding/base64"
import "fmt"
import "os"

const (
  _ = iota
  VERIFICATION_FAILED
  MISSING_HASHES
)

func main() {
  var m *MerkleTree

  k := os.Getenv("HMAC_KEY")
  p := os.Args[1:]

  m = Root(k, "+",
        Root(k, "*",
          Root(k, "3", nil, nil),
          Root(k, "2", nil, nil)),
        Root(k, "1", nil, nil))

  m.Each(func(m MerkleTree) {
    if len(p) == 0 {
      fmt.Println("Signature Verification Failed")
      os.Exit(MISSING_HASHES)
    } else {
      m.IfNodeInvalid(k, p[0], func() {
        fmt.Println("Signature Verification Failed")
        fmt.Printf("%v != %v\n", m.h, p[0])
        os.Exit(VERIFICATION_FAILED) })
      p = p[1:]
    }})
  fmt.Println("Signature Verification Succeeded")
}

type MerkleTree struct {
  V string
  h string
  l *MerkleTree
  r *MerkleTree
}

func Root(k, v string, l, r *MerkleTree) (m *MerkleTree) {
  m = &MerkleTree{ v, "", l, r }
  m.h = m.hash(k)
  return
}

func (m *MerkleTree) hash(k string) string {
  h := hmac.New(sha512.New, []byte(k))
  h.Write([]byte(m.V))
  if m.l != nil {
    h.Write([]byte(m.l.h))
  }
  if m.r != nil {
    h.Write([]byte(m.r.h))
  }
  return EncodeToString(h.Sum(nil))
}

func (m *MerkleTree) Each(f func(MerkleTree)) {
  if m != nil {
    f(*m)
    m.l.Each(f)
    m.r.Each(f)
  }
  return
}

func (m *MerkleTree) IfNodeInvalid(k, h string, f func()) {
  if h != m.hash(k) {
    f()
  }
}

func EncodeToString(b []byte) string {
  return base64.StdEncoding.EncodeToString(b)
}
