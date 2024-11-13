package meta

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Meta struct {
	s *goquery.Selection
}

// ReadDocument reads HTML document from r.
func ReadDocument(r io.Reader) (*Meta, error) {

	m := new(Meta)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)

	// Unwrap the "head" section
	m.s = doc.Unwrap()

	return m, err
}

// FetchDocument downloads and reads HTML document from url.
//
// Uses HTTP Get to get the content of the URL.
func FetchDocument(url string) (*Meta, error) {

	// Download
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: %s", resp.Status)
	}

	return ReadDocument(resp.Body)
}

// Title returns the site title.
//
//	<title>...</title>
//
// If not found, returns an empty string ("").
func (m *Meta) Title() string {
	return m.s.FindMatcher(goquery.Single(TitleSelector)).Text()
}

/*
Description returns the site description.

	<meta name="description" content="...">

If not found, returns an empty string.
*/
func (m *Meta) Description() string {
	return m.s.FindMatcher(goquery.Single(DescriptionSelector)).AttrOr("content", "")
}

func (m *Meta) Icons() []Icon {

	var v []Icon

	m.s.Find(IconSelector).Each(func(i int, s *goquery.Selection) {

		href := s.AttrOr("href", "")
		if len(href) == 0 {
			return
		}

		v = append(v, Icon{Href: href, Type: s.AttrOr("type", ""), Sizes: s.AttrOr("sizes", "")})
	})

	return v
}

func (m *Meta) Keywords() string {
	return m.s.FindMatcher(goquery.Single(KeywordsSelector)).AttrOr("content", "")
}

func (m *Meta) Robots() []string {

	cont := m.s.FindMatcher(goquery.Single(RobotsSelector)).AttrOr("content", "")

	if len(cont) == 0 {
		return nil
	}

	return strings.Split(cont, ",")
}

func (m *Meta) Rating() []string {

	var v []string

	m.s.Find(RatingSelector).Each(func(i int, s *goquery.Selection) {

		cont := s.AttrOr("content", "")
		if len(cont) == 0 {
			return
		}

		v = append(v, cont)
	})

	return v
}

func (m *Meta) Canonical() string {
	return m.s.FindMatcher(goquery.Single(CanonicalSelector)).AttrOr("href", "")
}

func (m *Meta) Html() (string, error) {
	return m.s.Html()
}

func (m *Meta) Alternate() []Alternate {

	var v []Alternate

	m.s.Find(AlternateSelector).Each(func(i int, s *goquery.Selection) {

		hreflang := s.AttrOr("hreflang", "")
		if len(hreflang) == 0 {
			return
		}

		href := s.AttrOr("href", "")
		if len(href) == 0 {
			return
		}

		v = append(v, Alternate{Lang: hreflang, URL: href})
	})

	return v
}
