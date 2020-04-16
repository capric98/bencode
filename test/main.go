package main

import (
	"fmt"
	"unsafe"
)

const (
	kindMask = (1 << 5) - 1
)

// tflag values must be kept in sync with copies in:
//	cmd/compile/internal/gc/reflect.go
//	cmd/link/internal/ld/decodesym.go
//	runtime/type.go
type tflag uint8
type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype
type textOff int32 // offset from top of text section
// rtype must be kept in sync with src/runtime/type.go:/^type._type.
type rtype struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}
type slice struct {
	data     uintptr
	len, cap int
}

func main() {
	d := "test stringtest stringtest stringtest stringtest stringtest string"
	var e = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(&d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(&e))
	ee := e
	ee[2] = 199
	fmt.Println(e)
}

func toByte(obj interface{}) {
	if dummy.b {
		dummy.x = obj
	} // escapes

	emp := (*emptyInterface)(unsafe.Pointer(&obj))

	fmt.Printf("%T object: %v\n", obj, obj)
	fmt.Printf(" - occupies %v bytes\n", emp.typ.size)
	fmt.Println(" - ptrdata", emp.typ.ptrdata)
	fmt.Printf("%#v\n", emp.typ)
	if int(emp.typ.ptrdata) != 0 {
		sl := len(*(*string)(emp.word))
		s := &slice{
			data: *(*uintptr)(emp.word),
			len:  sl,
			cap:  sl,
		}
		bs := *(*[]byte)(unsafe.Pointer(s))
		fmt.Println(bs)
	} else {
		sl := int(emp.typ.size)
		s := &slice{
			data: *(*uintptr)(emp.word),
			len:  sl,
			cap:  sl,
		}
		bs := *(*[]byte)(unsafe.Pointer(s))
		fmt.Println(bs)
	}
	fmt.Println("========")
}

var dummy struct {
	b bool
	x interface{}
}

func kind(k uint8) uint {
	return uint(k & kindMask)
}
