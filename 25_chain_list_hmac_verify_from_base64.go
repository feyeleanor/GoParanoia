package main

import "crypto/hmac"
import "crypto/sha512"
import "fmt"
import "os"

const (
	_ = iota
	VERIFICATION_FAILED
	MISMATCHED_PARAMETERS
)

func main() {
	var l *List
	var s *SignedList

	if len(os.Args)%2 != 1 {
		os.Exit(MISMATCHED_PARAMETERS)
	}

	k := os.Getenv("HMAC_KEY")
	for i := len(os.Args) - 1; i > 1; i-- {
		l = &List{os.Args[i], l}
		i--
		s = s.Push(k, os.Args[i])
	}

	s.Each(
		func(s SignedList) {
			l = l.Step(func(v string) {
				s.IfNodeInvalid(v, func() {
					fmt.Println("Signature Verification Failed")
					fmt.Printf("%v != %v\n", s.H, v)
					os.Exit(VERIFICATION_FAILED)
				})
			})
		})
	fmt.Println("Signature Verification Succeeded")
}

type List struct {
	V string
	*List
}

func (l *List) Step(f func(string)) (r *List) {
	if l != nil {
		f(l.V)
		r = l.List
	}
	return
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
		v,
		EncodeToString(h.Sum(nil)),
		s,
	}
}

func (s *SignedList) Each(f func(SignedList)) {
	if s != nil {
		f(*s)
		s.SignedList.Each(f)
	}
	return
}

func (s SignedList) IfNodeInvalid(h string, f func()) {
	if h != s.H {
		f()
	}
}
