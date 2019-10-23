package main

import "os"

const (
	_ = iota
	FILE_CREATE_FAILED
)

func main() {
	_, e := os.Create(os.Args[1])
	ExitOnError(e, FILE_CREATE_FAILED)
}
