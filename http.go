package main

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "os"
import "strings"

const DEFAULT_ADDRESS = ":3000"
const HTTP_ADDRESS = "HTTP_ADDRESS"
const OCTET_STREAM = "application/octet-stream"

func HTTP_readbody(r io.ReadCloser)  (s string) {
  defer r.Close()
  if b, e := ioutil.ReadAll(r); e == nil {
    s = string(b)
  }
  return
}

func HTTP_put(url, m string) (*http.Response, error) {
  return HTTP_doRequest("PUT", url, m)
}

func HTTP_delete(url, m string) (*http.Response, error) {
  return HTTP_doRequest("DELETE", url, m)
}

func HTTP_doRequest(method, url, m string) (r *http.Response, e error) {
  var req *http.Request

	req, e = http.NewRequest(method, url, strings.NewReader(m))
  if e == nil {
  	req.ContentLength = int64(len(m))
    r, e = http.DefaultClient.Do(req)
  }
	return
}

func HTTP_URL(p ...string) string {
  return "http://" + strings.Join(p, "/")
}

func HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
  http.HandleFunc(Route(p), f)
}

func SubPath(p string, r *http.Request) string {
  return r.URL.Path[len(Route(p)):]
}

func Route(p ...string) string {
  return fmt.Sprintf("/%v/", strings.Join(p, "/"))
}

func ServerAddress(a string) (r string) {
  r = os.Getenv(a)
	if r == "" {
    r = DEFAULT_ADDRESS
  }
  return
}
