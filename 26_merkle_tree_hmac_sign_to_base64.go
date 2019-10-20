package main

import "crypto/hmac"
import "crypto/sha512"
import "fmt"
import "os"

func main() {
	var m *MerkleTree

	k := os.Getenv("HMAC_KEY")

	m = Root(k, "+",
		Root(k, "*",
			Root(k, "3", nil, nil),
			Root(k, "2", nil, nil)),
		Root(k, "1", nil, nil))

	m.Each(
		func(m MerkleTree) {
			fmt.Println(string(m.h), m.V)
		})
}

type MerkleTree struct {
	V string
	h string
	l *MerkleTree
	r *MerkleTree
}

func Root(k, v string, l, r *MerkleTree) (m *MerkleTree) {
	m = &MerkleTree{v, "", l, r}
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
