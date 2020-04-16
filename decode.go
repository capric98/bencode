package bencode

import (
	"io"
	"io/ioutil"
	"unsafe"
)

// A Decoder reads and decodes B values from an input stream.
type Decoder struct {
	r           io.Reader
	buf         []byte
	pos, bufLen int

	b *B
}

// Decode decodes B values from an input []byte.
func Decode(body []byte) (*B, error) {
	cbuf := make([]byte, len(body))
	copy(cbuf, body)
	return (&Decoder{
		r:   nil,
		buf: cbuf,
		b:   &B{b: &bstruct{}},
	}).Decode()
}

// NewDecoder news a Decoder from given io.Reader.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:   r,
		buf: nil,
		b:   &B{b: &bstruct{}},
	}
}

// Decode decodes B values from *Decoder.
func (d *Decoder) Decode() (b *B, e error) {
	if d.r != nil {
		body, e := ioutil.ReadAll(d.r)
		if e != nil {
			return nil, e
		}
		d.buf = body
	}
	d.bufLen = len(d.buf)

	e = d.decode(d.b.b, 1)
	b = d.b
	return
}

func (d *Decoder) decode(bs *bstruct, depth byte) (e error) {
	if depth == 0 {
		return ErrMaxDepth
	}

	if d.pos >= d.bufLen {
		return io.EOF
	}

	// fmt.Println("Decode:", string(d.buf[d.pos:]))

	switch d.buf[d.pos] {
	case 'd':
		bs.typ = DICT
		d.pos++
		for {
			var key string
			d.decodeStr(&key)
			nbs := &bstruct{}
			e = d.decode(nbs, depth+1)
			if e != nil {
				return
			}
			if nbs.typ == NULL {
				return
			}
			bs.list = append(bs.list, pair{key: &key, value: nbs})
		}
	case 'l':
		bs.typ = LIST
		d.pos++
		for {
			nbs := &bstruct{}
			e = d.decode(nbs, depth+1)
			if e != nil {
				return
			}
			if nbs.typ == NULL {
				break
			}
			bs.list = append(bs.list, pair{value: nbs})
		}
	case 'e':
		d.pos++
		return
	case 'i':
		bs.typ = INT
		bs.n = d.decodeInt()
	default:
		bs.typ = STRING
		d.decodeStr(&bs.str)
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

func (d *Decoder) decodeStr(str *string) {
	pos := d.pos
	if pos >= d.bufLen {
		return
	}
	if d.buf[pos] == 'e' {
		return
	}
	// fmt.Println("decode string", string(d.buf[d.pos:]))

	slen := 0
	for ; pos < d.bufLen && d.buf[pos] != ':'; pos++ {
		slen = slen*10 + int(d.buf[pos]-'0')
	}
	startPos, endPos := pos+1, pos+slen+1
	if endPos > d.bufLen {
		return
	}

	d.pos = endPos
	sb := d.buf[startPos:endPos]
	*str = *(*string)(unsafe.Pointer(&sb))
}
