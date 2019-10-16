package main

import "fmt"
import "os"

const (
  _ = iota
  FILE_CREATE_FAILED
)

func main() {
  _, e := os.Create(os.Args[1])
  ExitOnError(e, FILE_CREATE_FAILED)
}

func ExitOnError(e error, n int) {
  if e != nil {
    fmt.Println(e)
    os.Exit(n)
  }
}
