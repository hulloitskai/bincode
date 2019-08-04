package encode

import (
	"io"
	"strings"

	"github.com/cockroachdb/errors"
)

// NewEncoder creates a new Encoder that encodes data to a binary representation
// format specified by key.
//
// The key must contain only alphabetical characters.
func NewEncoder(key string, w io.Writer) (io.Writer, error) {
	for i, c := range []byte(key) {
		if isNotAlpha(c) {
			return nil, errors.Newf("encode: key contains non-alphabetical "+
				"character at index %d", i)
		}
	}
	return Encoder{
		key: strings.ToLower(key),
		dst: w,
	}, nil
}

// An Encoder can encode data to a binary representation format.
type Encoder struct {
	key string
	dst io.Writer
}

func (e Encoder) Write(p []byte) (n int, err error) {
	var ki int // key index
	for _, b := range p {
		word := make([]byte, 8)

		for i := range word {
			c := e.key[ki]
			if (b>>uint(7-i))&1 == 1 { // black magic
				c = c - ('a' - 'A') // capitalize
			}
			ki = (ki + 1) % len(e.key)
			word[i] = c
		}

		if _, err := e.dst.Write(word); err != nil {
			return n, err
		}
		n++
	}
	return n, nil
}
