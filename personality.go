package main

import "fmt"

const BOB Person = "Bob"
const ALICE Person = "Alice"

type Person string

func (p Person) Report(m ...interface{}) {
  i := []interface{} { p }
  fmt.Println(append(i, m...)...)
}

func (p Person) ShowCurrentKeys(a *AES_channel) {
  p.Report("encodes messages with:", EncodeToString([]byte(a.ko)))
  p.Report("decodes messages with:", EncodeToString([]byte(a.ki)))
}
