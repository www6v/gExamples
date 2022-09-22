package oo

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

///// 30｜接口：Go中最强大的魔法 - Tony Bai
///// 包装器模式
func TestDecorator(t *testing.T) {
	r := strings.NewReader("hello, gopher!\n")
	lr := io.LimitReader(r, 4)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

/////
func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

type capitalizedReader struct {
	r io.Reader
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return 0, err
	}
	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func TestDecorator1(t *testing.T) {
	r := strings.NewReader("hello, gopher!\n")
	r1 := CapReader(io.LimitReader(r, 4))
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
