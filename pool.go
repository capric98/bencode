package bencode

import "sync"

var (
	pool = sync.Pool{
		New: func() interface{} {
			bs := make([]bstruct, 1, 32)
			return &bs
		},
	}
)
