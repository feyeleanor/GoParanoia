package main

import "fmt"
import "net/http"
import "os"
import "strings"

const DEFAULT_ADDRESS = ":3000"
const PING_URL = "ping"

const HTTP_ADDRESS = "HTTP_ADDRESS"

func init() {
	HandleFunc(PING_URL, func(w http.ResponseWriter, r *http.Request) {
    m := SubPath(PING_URL, r.URL.Path)
    fmt.Println("B:", m)
    fmt.Fprint(w, m)
  })
}

func main() {
  http.ListenAndServe(
    ServerAddress(HTTP_ADDRESS), nil)
}

func ServerAddress(a string) (r string) {
  r = os.Getenv(a)
	if r == "" {
    r = DEFAULT_ADDRESS
  }
  return
}

func SubPath(p, url string) string {
  return url[len(Route(p)):]
}

func HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
  http.HandleFunc(Route(p), f)
}

func Route(p ...string) string {
  return fmt.Sprintf("/%v/", strings.Join(p, "/"))
}
