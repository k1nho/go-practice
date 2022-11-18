package main

import "fmt"

type MyReader struct{}


func (r MyReader) Read(b []byte) (int, error) {
  n := len(b)

  for i := 0; i < n; i++ {
    b[i] = 65
  }

  return n, nil
}


func main() {
  var r MyReader

  if val, err := r.Read(make([]byte, 8)); err != nil {
    panic("could not read!")
  } else {
    fmt.Printf("Bytes read into buffer: %d\n", val)
  }
  
}
