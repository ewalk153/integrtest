package integrtest

import (
	"io"
	"io/ioutil"
)

func Foo(r io.Reader) (b []byte, err error) {
	b, err = ioutil.ReadAll(r)
	return
}
