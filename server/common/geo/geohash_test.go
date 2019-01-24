package geo

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	EncodeTest()
}

func TestIndexBound(t *testing.T) {
	code := EncodeInt(0, 123)
	l, u := GetIndexBound(code)
	fmt.Println(code, l, u)
	fmt.Printf("%b %b %b\n", code, l, u)
}
