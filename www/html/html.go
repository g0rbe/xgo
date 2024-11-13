package html

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

// ReadBody reads HTML document from r and unwraps the body.
func ReadBody(r io.Reader) ([]byte, error) {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)

	// Unwrap the "head" section
	body, err := doc.FilterMatcher(goquery.Single("body")).Html()
	if err != nil {
		return nil, err
	}

	return []byte(body), nil
}
