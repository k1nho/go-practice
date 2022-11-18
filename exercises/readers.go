package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyReader struct{}


func (r MyReader) Read(b []byte) (int, error) {
  n := len(b)

  for i := 0; i < n; i++ {
    b[i] = 65
  }

  return n, nil
}

type rot13Reader struct {
	r io.Reader
}

func rotateByte(b byte) byte {
	switch {
		case b >= 97 && b <= 109:
			b += 13
		case b >= 110 && b <= 122:
			b -=13
	}
	return b
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i:= 0; i <= n; i++ { 
		b[i] = rotateByte(b[i])
	}
	return n, nil
}



func main() {
  var r MyReader

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rotReader := rot13Reader{s}
	io.Copy(os.Stdout, &rotReader)

  if val, err := r.Read(make([]byte, 8)); err != nil {
    panic("could not read!")
  } else {
    fmt.Printf("Bytes read into buffer: %d\n", val)
  }
  
}


