package main

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "os"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const PING_URL = "/ping/"

const HTTP_ADDRESS = "HTTP_ADDRESS"
const NO_BODY_RECEIVED = "HTTP message has no body"

func init() {
	http.HandleFunc(PING_URL, func(w http.ResponseWriter, r *http.Request) {
    m := r.URL.Path[len(PING_URL):]
    fmt.Println("B:", m)
    fmt.Fprint(w, m)
  })
}

func main() {
  a := os.Getenv(HTTP_ADDRESS)
	if a == "" {
    a = DEFAULT_ADDRESS
  }

  go func() {
    http.ListenAndServe(a, nil)
  }()

  url := HTTP + a + PING_URL + "%v"
  for _, v := range os.Args[1:] {
    r, e := http.Get(fmt.Sprintf(url, v))
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
