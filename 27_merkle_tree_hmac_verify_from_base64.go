package main

import "fmt"
import "os"

func main() {
	k := os.Getenv("HMAC_KEY")
	m := Tree(k, "+",
		Tree(k, "*",
			Tree(k, "-",
        Tree(k, "x"),
        Tree(k, "y")),
			Tree(k, "2"),
			Tree(k, "7")),
		Tree(k, "1"))

	p := os.Args[1:]
	m.Each(func(m MerkleTree) {
		if len(p) == 0 {
			fmt.Println("Signature Verification Failed")
			os.Exit(MISSING_HASHES)
		} else {
      if p[0] != m.hash(k) {
				fmt.Println("Signature Verification Failed")
				fmt.Println(m.h, "!=", p[0])
				os.Exit(VERIFICATION_FAILED)
			}
			p = p[1:]
		}
	})
	fmt.Println("Signature Verification Succeeded")
}

type MerkleTree struct {
	V string
	h string
  t []*MerkleTree
}

func Tree(k, v string, t ...*MerkleTree) (m *MerkleTree) {
	m = &MerkleTree{ v, "", make([]*MerkleTree, len(t)) }
  copy(m.t, t)
	m.h = m.hash(k)
	return
}

func (m *MerkleTree) hash(k string) string {
  h := []string{ m.V }
  for _, v := range m.t {
    h = append(h, v.h)
  }
	return EncodeToString(HMAC_SignAll(k, h...))
}

func (m *MerkleTree) Each(f func(MerkleTree)) {
	if m != nil {
		f(*m)
    for _, v := range m.t {
      v.Each(f)
    }
	}
	return
}
