package main

import "fmt"

func main() {
  fmt.Println(TrimBuffer([]byte{}))

  fmt.Println(TrimBuffer([]byte{ 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 0, 0 }))

  fmt.Println(TrimBuffer([]byte{ 1 }))
  fmt.Println(TrimBuffer([]byte{ 0, 1 }))
  fmt.Println(TrimBuffer([]byte{ 0, 0, 1 }))

  fmt.Println(TrimBuffer([]byte{ 1, 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 1, 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 0, 1, 0 }))

  fmt.Println(TrimBuffer([]byte{ 1, 0, 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 1, 0, 0 }))
  fmt.Println(TrimBuffer([]byte{ 0, 0, 1, 0, 0 }))
}

func TrimBuffer(m []byte) (b []byte) {
  i := len(m)
  for ; i > 0 && m[i - 1] == 0; i-- {}
	return m[:i]
}
