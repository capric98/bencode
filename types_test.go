package bencode

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestTypeSize(t *testing.T) {
	fmt.Println(unsafe.Sizeof(B{}))  // 48
	fmt.Println(unsafe.Sizeof(kv{})) // 16
	t.Fail()
}
