// +build integration

package integrtest

import (
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	f, err := os.Open("data/data.txt")
	if err != nil {
		t.Error(err)
	}
	b, err := Foo(f)
	if err != nil {
		t.Error(err)
	}
	if len(b) == 0 {
		t.Error("Wanted bytes string with length it was 0")
	}
}
