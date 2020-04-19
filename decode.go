package bencode

import (
	"io"
	"io/ioutil"
	"unsafe"
)

// A Decoder reads and decodes B values from an input stream.
type Decoder struct {
	r      io.Reader
	buf    []byte
	bufLen int
	pos    int

	result B
}

// Decode decodes B values from an input []byte.
func Decode(body []byte) (B, error) {
	return (&Decoder{
		r:   nil,
		buf: append(body[:0:0], body...),
		//b:   &B{b: &bstruct{}},
	}).Decode()
}

// NewDecoder news a Decoder from given io.Reader.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:   r,
		buf: nil,
		//b:   &B{b: &bstruct{}},
	}
}

// Decode decodes B values from *Decoder.
func (d *Decoder) Decode() (b B, e error) {
	if d.r != nil {
		body, e := ioutil.ReadAll(d.r)
		if e != nil {
			return b, e
		}
		d.buf = body
	}
	d.bufLen = len(d.buf)
	d.result.b = (pool.Get()).(*[]bstruct)

	e = d.decode(0, 1)
	if e == nil {
		cb := append((*(d.result.b))[:0:0], (*(d.result.b))...)
		b = B{
			pos: 0,
			b:   &cb,
		}
	}
	go func() {
		ob := *(d.result.b)
		for i := range ob {
			ob[i].typ = NULL
			ob[i].keys = ob[i].keys[:0]
			ob[i].list = ob[i].list[:0]
		}
		ob = ob[:1]
		pool.Put(&ob)
	}()

	return
}

func (d *Decoder) decode(bpos int, depth byte) (e error) {
	if depth == 0 {
		return ErrMaxDepth
	}

	if d.pos >= d.bufLen {
		return io.EOF
	}

	switch d.buf[d.pos] {
	case 'd':
		(*(d.result.b))[bpos].typ = DICT
		d.pos++
		for {
			nbpos := d.result.next()

			key := d.decodeStr()
			e = d.decode(nbpos, depth+1)
			if e != nil {
				return
			}
			if key != nil && (*(d.result.b))[nbpos].typ != NULL {
				(*(d.result.b))[bpos].keys = append((*(d.result.b))[bpos].keys, *key)
				(*(d.result.b))[bpos].list = append((*(d.result.b))[bpos].list, nbpos)
			} else {
				d.result.retrive(nbpos)
				return
			}
		}
	case 'l':
		(*(d.result.b))[bpos].typ = LIST
		d.pos++
		for {
			nbpos := d.result.next()

			e = d.decode(nbpos, depth+1)
			if e != nil {
				return
			}
			if (*(d.result.b))[nbpos].typ != NULL {
				(*(d.result.b))[bpos].list = append((*(d.result.b))[bpos].list, nbpos)
			} else {
				d.result.retrive(nbpos)
				return
			}
		}
	case 'e':
		d.pos++
		(*(d.result.b))[bpos].typ = NULL
		(*(d.result.b))[bpos].keys = nil
		(*(d.result.b))[bpos].list = nil
		return
	case 'i':
		(*(d.result.b))[bpos].typ = INT
		(*(d.result.b))[bpos].n = d.decodeInt()
	default:
		(*(d.result.b))[bpos].typ = STRING
		if key := d.decodeStr(); key != nil {
			(*(d.result.b))[bpos].str = *key
			// (*(d.result.b))[bpos].keys = append((*(d.result.b))[bpos].keys, *key)
		}
	}

	return
}

func (d *Decoder) decodeInt() int64 {
	pos := d.pos
	if pos >= d.bufLen {
		return 0
	}
	if d.buf[pos] != 'i' {
		return 0
	}
	pos++

	value := int64(0)
	positive := true
	if d.buf[pos] == '-' {
		positive = false
		pos++
	}

	for ; pos < d.bufLen && d.buf[pos] != 'e'; pos++ {
		value = value*10 + int64(d.buf[pos]-'0')
	}
	d.pos = pos + 1

	if positive {
		return value
	}
	return -value
}

func (d *Decoder) decodeStr() *string {
	pos := d.pos
	if pos >= d.bufLen {
		return nil
	}
	if d.buf[pos] == 'e' {
		return nil
	}
	// fmt.Println("decode string", string(d.buf[d.pos:]))

	slen := 0
	for ; pos < d.bufLen && d.buf[pos] != ':'; pos++ {
		slen = slen*10 + int(d.buf[pos]-'0')
	}
	startPos, endPos := pos+1, pos+slen+1
	if endPos > d.bufLen {
		return nil
	}

	d.pos = endPos
	bstr := d.buf[startPos:endPos]
	return (*string)(unsafe.Pointer(&bstr))
}

func (b B) retrive(pos int) {
	if pos < len(*(b.b)) {
		(*(b.b)) = (*(b.b))[:pos]
	}
}

func (b B) next() int {
	*(b.b) = append(*(b.b), bstruct{})
	return len(*(b.b)) - 1
}
