package bencode

import (
	"io"
	"io/ioutil"
)

// A Decoder reads and decodes B values from an input stream.
type Decoder struct {
	r      io.Reader
	buf    []byte
	bufLen int
	pos    int
}

// Decode decodes B values from an input []byte.
func Decode(body []byte) (B, error) {
	return (&Decoder{
		r: nil,
		//buf: append(body[:0:0], body...),
		buf: body,
	}).Decode()
}

// NewDecoder news a Decoder from given io.Reader.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:   r,
		buf: nil,
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

	d.decode(&b, 1)

	return
}

// func (d *Decoder) decode(bpos int, depth byte) (e error) {
// 	if depth == 0 {
// 		return ErrMaxDepth
// 	}

// 	if d.pos >= d.bufLen {
// 		return io.EOF
// 	}

// 	switch d.buf[d.pos] {
// 	case 'd':
// 		(*(d.result.b))[bpos].typ = DICT
// 		d.pos++
// 		for {
// 			nbpos := d.result.next()

// 			key := d.decodeStr()
// 			e = d.decode(nbpos, depth+1)
// 			if e != nil {
// 				return
// 			}
// 			if key != nil && (*(d.result.b))[nbpos].typ != NULL {
// 				(*(d.result.b))[bpos].keys = append((*(d.result.b))[bpos].keys, *key)
// 				(*(d.result.b))[bpos].list = append((*(d.result.b))[bpos].list, nbpos)
// 			} else {
// 				d.result.retrive(nbpos)
// 				return
// 			}
// 		}
// 	case 'l':
// 		(*(d.result.b))[bpos].typ = LIST
// 		d.pos++
// 		for {
// 			nbpos := d.result.next()

// 			e = d.decode(nbpos, depth+1)
// 			if e != nil {
// 				return
// 			}
// 			if (*(d.result.b))[nbpos].typ != NULL {
// 				(*(d.result.b))[bpos].list = append((*(d.result.b))[bpos].list, nbpos)
// 			} else {
// 				d.result.retrive(nbpos)
// 				return
// 			}
// 		}
// 	case 'e':
// 		d.pos++
// 		(*(d.result.b))[bpos].typ = NULL
// 		(*(d.result.b))[bpos].keys = nil
// 		(*(d.result.b))[bpos].list = nil
// 		return
// 	case 'i':
// 		(*(d.result.b))[bpos].typ = INT
// 		(*(d.result.b))[bpos].n = d.decodeInt()
// 	default:
// 		(*(d.result.b))[bpos].typ = STRING
// 		if key := d.decodeStr(); key != nil {
// 			(*(d.result.b))[bpos].str = *key
// 			// (*(d.result.b))[bpos].keys = append((*(d.result.b))[bpos].keys, *key)
// 		}
// 	}

// 	return
// }

func (d *Decoder) decode(target *B, depth byte) {
	if depth == 0 {
		panic(ErrMaxDepth)
	}

	switch d.buf[d.pos] {
	case 'd':
		d.pos++
		target.typ = DICT
		for {
			tmp := kv{key: d.decodeStr(), value: new(B)}
			d.decode(tmp.value, depth+1)
			if tmp.value.typ == NULL {
				return
			}
			target.list = append(target.list, tmp)
		}
	case 'l':
		d.pos++
		target.typ = LIST
		for {
			tmp := kv{key: nil, value: new(B)}
			d.decode(tmp.value, depth+1)
			if tmp.value.typ == NULL {
				return
			}
			target.list = append(target.list, tmp)
		}
	case 'e':
		d.pos++
		target.typ = NULL
		return
	case 'i':
		target.typ = INT
		target.n = d.decodeInt()
	default:
		target.typ = STRING
		target.str = d.decodeStr()
	}

	return
}

func (d *Decoder) decodeInt() int64 {
	pos := d.pos
	if d.buf[pos] != 'i' {
		panic("i")
	}
	pos++

	positive := true
	if d.buf[pos] == '-' {
		positive = false
		pos++
	}

	value := int64(0)
	for ; d.buf[pos] != 'e'; pos++ {
		value = value*10 + int64(d.buf[pos]-'0')
	}
	d.pos = pos + 1

	if positive {
		return value
	}
	return -value
}

func (d *Decoder) decodeStr() []byte {
	pos := d.pos
	if d.buf[pos] == 'e' {
		return nil
	}

	slen := 0
	for ; d.buf[pos] != ':'; pos++ {
		slen = slen*10 + int(d.buf[pos]-'0')
	}
	endPos := pos + slen + 1
	if endPos > d.bufLen {
		return nil
	}

	d.pos = endPos
	return d.buf[pos+1 : endPos]
}
