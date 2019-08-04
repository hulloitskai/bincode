package encode

import (
	"io"
)

// AlphaFilter creates an io.Reader that filters out non-alphabet characters.
func AlphaFilter(r io.Reader) io.Reader {
	return alphaFilter{r}
}

// An alphaFilter is an io.Reader that filters out non-alphabet characters.
type alphaFilter struct{ src io.Reader }

func (af alphaFilter) Read(p []byte) (n int, err error) {
	rem := len(p)

Read:
	tmp := make([]byte, rem)
	if _, err := af.src.Read(tmp); err != nil {
		return len(p) - rem, err
	}

	r := rem
	for i := 0; i < r; i++ {
		b := tmp[i]
		if isNotAlpha(b) {
			continue
		}
		p[len(p)-rem] = b
		rem--
	}

	if rem > 0 {
		goto Read
	}
	return len(p), nil
}

func isNotAlpha(b byte) bool {
	return b < 'A' || 'z' < b || ('Z' > b && b > 'a')
}
