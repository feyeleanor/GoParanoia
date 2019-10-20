package main

import "encoding/base64"
import "fmt"
import "os"

func main() {
	k := os.Getenv("AES_KEY")
	if m, e := AESEncrypt(k, os.Args[1]); e == nil {
		PrintEncrypted(m)
	} else {
		fmt.Printf("error: %v\n", e)
	}
}

func PrintEncrypted(m []byte) {
	fmt.Println(base64.StdEncoding.EncodeToString(m))
}
