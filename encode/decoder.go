package encode

import (
	stderrs "errors"
	"io"

	"github.com/cockroachdb/errors"
)

// NewDecoder creates a new Decoder that decodes data from a binary
// representation format.
func NewDecoder(r io.Reader, opts ...func(*DecoderConfig)) io.Reader {
	cfg := DecoderConfig{
		SanitizeInputs: true,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.SanitizeInputs {
		r = AlphaFilter(r)
	}
	return Decoder{src: r}
}

type (
	// A Decoder can decode data from a binary representation format.
	Decoder struct {
		src io.Reader
	}

	// A DecoderConfig configures a Decoder.
	DecoderConfig struct {
		// SanitizeInputs controls whether or not the Decoder sanitizes inputs
		// by stripping non-alphabetical characters.
		SanitizeInputs bool
	}
)

func (d Decoder) Read(p []byte) (n int, err error) {
	for i := range p {
		word := make([]byte, 8)
		n, err := d.src.Read(word)
		if n != len(word) {
			if err != nil {
				if errors.Is(err, io.EOF) && (n > 0) {
					return i, ErrShortWord
				}
				return i, err
			}
			return i, ErrShortWord
		}

		var b byte
		for j, c := range word {
			if j > 0 {
				b <<= 1
			}
			if c <= 'Z' { // uppercase
				b++
			}
		}
		p[i] = b

		if err != nil {
			return i + 1, err
		}
	}
	return len(p), nil
}

// ErrShortWord is returned by a Decoder when it attempts to decode a binary
// 'word', only to receive less than 8 bits from the source reader.
var ErrShortWord = stderrs.New("encode: word too short")
