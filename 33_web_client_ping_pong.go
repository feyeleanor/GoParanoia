package main

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "os"

const DEFAULT_ADDRESS = ":3000"
const PING_URL = "/ping/"

const HTTP_ADDRESS = "HTTP_ADDRESS"
const NO_BODY_RECEIVED = "HTTP message has no body"

func main() {
  a := ServerAddress(HTTP_ADDRESS)
  for _, v := range os.Args[1:] {
    r, e := http.Get(ServerURL(a, PING_URL, v))
    ExitOnError(e, WEB_REQUEST_FAILED)

    s := read_body(r.Body)
    if s == "" {
      s = NO_BODY_RECEIVED
    }
    fmt.Println("A:", s)
  }
}

func read_body(r io.ReadCloser) (s string) {
  defer r.Close()
  if b, e := ioutil.ReadAll(r); e == nil {
    s = string(b)
  }
  return
}

func ServerURL(a, p, v string) string {
  return "http://" + a + PING_URL + v
}

func ServerAddress(a string) (r string) {
  r = os.Getenv(a)
	if r == "" {
    r = DEFAULT_ADDRESS
  }
  return
}
