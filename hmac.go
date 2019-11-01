package main

import "crypto/hmac"
import "crypto/sha512"

func HMAC_Sign(k, m string) []byte {
	h := hmac.New(sha512.New, []byte(k))
	h.Write([]byte(m))
	return h.Sum(nil)
}

func HMAC_Verify(k, hs, m string) bool {
	return hmac.Equal(
		[]byte(hs),
		HMAC_Sign(k, m))
}
