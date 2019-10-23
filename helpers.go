package main

import "encoding/base64"
import "fmt"
import "os"

func EncodeToString(m []byte) string {
	return base64.StdEncoding.EncodeToString(m)
}

func DecodeString(s string) (b []byte, e error) {
	return base64.StdEncoding.DecodeString(s)
}

func ExitOnError(e error, n int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(n)
	}
}

func read_base64(s string) string {
	b, _ := DecodeString(s)
	return string(b)
}