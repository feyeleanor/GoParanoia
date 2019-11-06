package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"

const DEFAULT_ADDRESS = ":3000"
const HTTP = "http://"
const PING = "/ping/"

func init() {
	http.HandleFunc("/ping/", func(w http.ResponseWriter, r *http.Request) {
    m := r.URL.Path[len("/ping/"):]
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

  url := HTTP + a + PING + "%v"
  for _, v := range os.Args[1:] {
    r, e := http.Get(fmt.Sprintf(url, v))
    ExitOnError(e, WEB_REQUEST_FAILED)

    defer r.Body.Close()
    var b []byte
    b, e = ioutil.ReadAll(r.Body)
    ExitOnError(e, WEB_NO_BODY)
    fmt.Println("A:", string(b))
  }
}
