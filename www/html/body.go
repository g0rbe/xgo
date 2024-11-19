package html

import (
	"bytes"
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/g0rbe/xgo/cryptography/checksum"
)

var (
	ErrNoBody = errors.New("body not foound")
)

type Body struct {
	s *goquery.Selection
}

// ParseBody reads HTML document from b bytes and unwraps the body.
func ParseBody(b []byte) (*Body, error) {

	buf := bytes.NewReader(b)
	return ReadBody(buf)
}

// ReadBody reads HTML document from r and unwraps the body.
//
// If the body not found, returns ErrNoBody.
func ReadBody(r io.Reader) (*Body, error) {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	// Unwrap the "body" section
	bodySelection := doc.Find("body").Unwrap()

	if bodySelection.Length() == 0 {
		return nil, ErrNoBody
	}

	v := new(Body)
	v.s = bodySelection

	return v, nil
}

func (b *Body) HTML() ([]byte, error) {
	v, err := b.s.Html()
	return []byte(v), err
}

func (b *Body) SHA256() checksum.SHA256 {
	v, err := b.HTML()
	if err != nil {
		return checksum.SHA256{}
	}

	return checksum.Data256(v)
}
