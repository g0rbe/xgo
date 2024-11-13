package sitemap

import (
	"encoding/xml"
	"net/url"
)

type Location url.URL

func ParseLocation(v string) (*Location, error) {

	u, err := url.Parse(v)
	if err != nil {
		return nil, err
	}

	return (*Location)(u), nil
}

func MustParseLocation(v string) *Location {

	u, err := url.Parse(v)
	if err != nil {
		panic(err)
	}

	return (*Location)(u)
}

func (l *Location) String() string {
	// Convert to *url.URL type
	return (*url.URL)(l).String()
}

// MarshalXML implements the xml.Marshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (l *Location) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "loc"

	return e.EncodeElement(l.String(), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (l *Location) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	u, err := url.Parse(v)
	if err != nil {
		return err
	}

	*l = (Location)(*u)

	return nil
}
