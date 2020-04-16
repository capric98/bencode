package bencode

import "io"

// An Encoder writes B values to an output stream.
type Encoder struct {
	w io.Writer
}
