package meta

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Meta struct {
	s *goquery.Selection
}

/*
FetchHeader downloads and returns the content of the <head>...</head> section.

Uses HTTP Get to get the content of the URL.
*/
func FetchHeader(url string) (*Meta, error) {

	// Download
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: %s", resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("document error: %w", err)
	}

	m := new(Meta)

	// Unwrap the "head" section
	m.s = doc.FindMatcher(goquery.Single("html head")).Unwrap()

	return m, err
}

/*
Title returns the site title.

	<title>...</title>

If not found, returns an empty string ("").
*/
func (m *Meta) Title() string {
	return m.s.FindMatcher(goquery.Single("title")).Text()
}

/*
Description returns the site description.

	<meta name="description" content="...">

If not found, returns an empty string.
*/
func (m *Meta) Description() string {
	return m.s.FindMatcher(goquery.Single("meta[name=\"description\" i][content]")).AttrOr("content", "")
}

func (m *Meta) Icons() []Icon {

	var v []Icon

	m.s.Find("link[rel=\"icon\" i], link[rel=\"shortcut icon\" i]").Each(func(i int, s *goquery.Selection) {

		href := s.AttrOr("href", "")
		if len(href) == 0 {
			return
		}

		v = append(v, Icon{Href: href, Type: s.AttrOr("type", ""), Sizes: s.AttrOr("sizes", "")})
	})

	return v
}

func (m *Meta) Keywords() []string {

	cont := m.s.FindMatcher(goquery.Single("meta[name=\"keywords\" i][content]")).AttrOr("content", "")

	if len(cont) == 0 {
		return nil
	}

	return strings.Split(cont, ", ")
}

func (m *Meta) Robots() []string {

	cont := m.s.FindMatcher(goquery.Single("meta[name=\"robots\" i][content]")).AttrOr("content", "")

	if len(cont) == 0 {
		return nil
	}

	return strings.Split(cont, ",")
}

func (m *Meta) Rating() []string {

	var v []string

	m.s.Find("meta[name=\"rating\" i][content]").Each(func(i int, s *goquery.Selection) {

		cont := s.AttrOr("content", "")
		if len(cont) == 0 {
			return
		}

		v = append(v, cont)
	})

	return v
}

func (m *Meta) Canonical() string {

	return m.s.FindMatcher(goquery.Single("link[rel=\"canonical\" i][href]")).AttrOr("href", "")

}

func (m *Meta) Html() (string, error) {
	return m.s.Html()
}
