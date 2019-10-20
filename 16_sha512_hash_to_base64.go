package main

import "fmt"
import "crypto/sha512"
import "os"

func main() {
	m := os.Args[1]
	h := sha512.New()
	h.Write([]byte(m))
	fmt.Println(EncodeToString(h.Sum(nil)))
}
