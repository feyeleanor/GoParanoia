package main

import "fmt"
import "os"

func main() {
  var s *SignedList
	var h string
  var ok bool

  a := os.Args[1:]
	if len(a)%2 != 0 {
		os.Exit(UNEVEN_PARAMETERS)
	}

	k := os.Getenv("HMAC_KEY")
  for i := len(a) - 1; i > 0; i-- {
    h = a[i]
    i--
    if s, ok = s.PushAndCheck(k, h, a[i]); !ok {
  		fmt.Println("Signature Verification Failed")
			fmt.Printf("%v != %v\n", h, s.H)
			os.Exit(VERIFICATION_FAILED)
    }
  }
	fmt.Println("Signature Verification Succeeded")
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
		b = HMAC_SignAll(k, "", v)
	} else {
		b = HMAC_SignAll(k, s.H, v)
	}
	return &SignedList{v, EncodeToString(b), s}
}

func (s *SignedList) PushAndCheck(k, h, v string) (r *SignedList, ok bool) {
  r = s.Push(k, v)
  ok = r.H == h
  return
}
