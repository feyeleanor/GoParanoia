package main

import "bytes"
import "encoding/base64"
import "fmt"
import "io"
import "io/ioutil"
import "os"

func EncodeToBase64(s string) string {
  return EncodeToString([]byte(s))
}

func EncodeToString(m []byte) string {
	return base64.StdEncoding.EncodeToString(m)
}

func DecodeString(s string) (b []byte, e error) {
	return base64.StdEncoding.DecodeString(s)
}

func EncodeStrings(b ...[]byte) (r string) {
	for _, v := range b {
		r += EncodeToString(v)
	}
	return
}

func EncodeToReader(m []byte) io.Reader {
	return bytes.NewBufferString(EncodeToString(m))
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

func LoadFile(s string) (b []byte) {
	var e error
	b, e = ioutil.ReadFile(s)
	ExitOnError(e, FILE_UNREADABLE)
	return
}
