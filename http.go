package main

import "io"
import "io/ioutil"

func HTTP_readbody(r io.ReadCloser) (s string, e error) {
  var b []byte

  defer r.Close()
  if b, e = ioutil.ReadAll(r); e == nil {
    s = string(b)
  }
  return
}
