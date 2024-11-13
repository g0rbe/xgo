package sitemap

import (
	"encoding/xml"
	"fmt"
)

// XHTMLNS is the xhtml namespace
const XHTMLNS = "http://www.w3.org/1999/xhtml"

// Alternate specify a language and region variant of the URL.
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
//
// See: https://developers.google.com/search/docs/specialty/international/localized-versions#sitemap
type Alternate struct {
	Language string // hreflang attribute
	URL      string // href attribute
}

// NewAlternate creates a new Alternate from href and hreflang.
func NewAlternate(hreflang, href string) Alternate {
	return Alternate{Language: hreflang, URL: href}
}

// MarshalXML implements the xml.Marshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (a Alternate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Set tag
	start.Name.Local = "xhtml:link"

	v := struct {
		Relationship string `xml:"rel,attr"`
		Language     string `xml:"hreflang,attr"`
		URL          string `xml:"href,attr"`
	}{
		Relationship: "alternate",
		Language:     a.Language,
		URL:          a.URL,
	}

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (a *Alternate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	if start.Name.Space == XHTMLNS {
		start.Name.Space = "xhtml"
	}

	v := struct {
		XMLName      xml.Name `xml:"xhtml link"`
		Relationship string   `xml:"rel,attr"`
		Language     string   `xml:"hreflang,attr"`
		URL          string   `xml:"href,attr"`
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	if v.Relationship != "alternate" {
		return fmt.Errorf("invalid \"rel\" attribute value: %s", v.Relationship)
	}

	if a == nil {
		a = new(Alternate)
	}

	a.Language = v.Language
	a.URL = v.URL

	return nil
}
