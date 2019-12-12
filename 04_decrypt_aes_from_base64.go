package main

import "encoding/base64"
import "fmt"
import "os"

func main() {
	k := os.Getenv("AES_KEY")
	s := ReadBase64(os.Args[1])
	if m, e := AES_DecryptAndTrim(k, s); e == nil {
		fmt.Println(m)
	} else {
		fmt.Printf("error: %v\n", m)
	}
}

func ReadBase64(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}
