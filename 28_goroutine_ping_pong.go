package main

import "fmt"
import "os"

func main() {
	AtoB := make(chan string)
  BtoA := Launch(AtoB, func(in, out chan string) {
    for m := range in {
      fmt.Println("B:", m)
      out <- m
    }
  })

  for _, v := range os.Args[1:] {
    AtoB <- v
    fmt.Println("A:", <- BtoA)
  }
}

func Launch(in chan string, f func(chan string, chan string)) (out chan string) {
  out = make(chan string)
  go func(i, o chan string) {
    f(i, o)
  }(in, out)
  return
}
