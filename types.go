package bencode

import "errors"

// Bencode types.
const (
	NULL = iota
	DICT
	LIST
	INT
	STRING

	MaxDepth = 100
)

// B illustrates a bencode struct.
type B struct {
	typ  int
	n    int64
	str  []byte
	list []kv
}

type kv struct {
	key   []byte
	value *B
}

// Some vars.
var (
	ErrMaxDepth            = errors.New("stack overflow")
	ErrDicWithNonStringKey = errors.New("dictionary's key must be a byte string")
)
