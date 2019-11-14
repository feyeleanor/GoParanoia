package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const PING_URL = "/ping/"

func init() {
	http.HandleFunc(PING_URL, func(w http.ResponseWriter, r *http.Request) {
    m := r.URL.Path[len(PING_URL):]
    fmt.Println("B:", m)
    fmt.Fprint(w, m)
  })
}

func main() {
  a := os.Getenv("HTTP_ADDRESS")
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

    fmt.Println("A:", HTTP_read(r))
  }
}

func HTTP_read(r *http.Response) string {
  defer r.Body.Close()
  b, e := ioutil.ReadAll(r.Body)
  ExitOnError(e, WEB_NO_BODY)
  return string(b)
}
