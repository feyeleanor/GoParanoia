package main

import "encoding/base64"
import "fmt"
import "os"

func main() {
	k := os.Getenv("AES_KEY")
	s := read_base64(os.Args[1])
	if m, e := AES_Decrypt(k, s); e == nil {
		fmt.Println(m)
	} else {
		fmt.Printf("error: %v\n", m)
	}
}

func read_base64(s string) string {
	b, _ := base64.URLEncoding.DecodeString(s)
	return string(b)
}
