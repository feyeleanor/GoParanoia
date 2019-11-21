package main

import "io"
import "io/ioutil"
import "net/http"
import "strings"

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

func HTTP_URL(a, p string, n ...string) string {
  return "http://" + a + p + strings.Join(n, "/")
}
