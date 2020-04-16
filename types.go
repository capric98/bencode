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
	b *bstruct
}

type bstruct struct {
	typ  byte
	n    int64
	str  string
	list []pair
}

type pair struct {
	key   *string
	value *bstruct
}

// Some vars.
var (
	ErrMaxDepth            = errors.New("stack overflow")
	ErrDicWithNonStringKey = errors.New("dictionary's key must be a byte string")
)
