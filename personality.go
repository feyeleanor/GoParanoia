package main

import "fmt"

type Person string

func (p Person) Report(m ...interface{}) {
  i := []interface{} { p }
  fmt.Println(append(i, m...)...)
}

func (p Person) ShowCurrentKeys(a *AES_channel) {
  p.Report("encodes messages with:", EncodeToString([]byte(a.ko)))
  p.Report("decodes messages with:", EncodeToString([]byte(a.ki)))
}
