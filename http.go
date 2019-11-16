package main

import "io"
import "io/ioutil"
import "net/http"
import "strings"

func HTTP_readbody(r io.ReadCloser) (s string, e error) {
  var b []byte

  defer r.Close()
  if b, e = ioutil.ReadAll(r); e == nil {
    s = string(b)
  }
  return
}

func HTTP_put(url, m string) (*http.Response, error) {
	r, e := http.NewRequest("PUT", url, strings.NewReader(m))
	ExitOnError(e, WEB_REQUEST_FAILED)
	r.ContentLength = int64(len(m))
	return http.DefaultClient.Do(r)
}
