package main

import "crypto/hmac"
import "crypto/sha512"
import "fmt"
import "os"

func main() {
	var s *SignedList

	k := os.Getenv("HMAC_KEY")
	for v := os.Args[1:]; len(v) > 0; v = v[1:] {
		s = s.Push(k, v[0])
	}
	s.Each(
		func(h, v string) {
			fmt.Println(h, v)
		})
}

type SignedList struct {
	V string
	H string
	*SignedList
}

func (s *SignedList) Each(f func(h, v string)) {
	if s != nil {
		f(s.H, s.V)
		s.SignedList.Each(f)
	}
	return
}

func (s *SignedList) Push(k, v string) *SignedList {
	var b []byte
	if s == nil {
		b = SignAll(k, "", v)
	} else {
		b = SignAll(k, s.H, v)
	}
	return &SignedList{v, EncodeToString(b), s}
}

func SignAll(k string, m ...string) []byte {
	h := hmac.New(sha512.New, []byte(k))
	for _, v := range m {
		h.Write([]byte(v))
	}
	return h.Sum(nil)
}
