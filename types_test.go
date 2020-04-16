package bencode

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestTypeSize(t *testing.T) {
	fmt.Println(unsafe.Sizeof(B{}))       // 8
	fmt.Println(unsafe.Sizeof(pair{}))    // 16
	fmt.Println(unsafe.Sizeof(bstruct{})) // 56
	t.Fail()
}

func TestOverflow(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Fail()
		}
	}()
	n := int64(1)
	for i := 0; i < 100; i++ {
		n = n * 2
	}
}
