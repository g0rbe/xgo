package sitemap

import (
	"crypto/sha256"
	"encoding/xml"
	"net/url"
)

// Location is the <loc> field of the sitemap.
//
//	<loc>https://example.com/sample1.html</loc>
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

// SHA256 create an SHA256 sum from Location.String().
//
// This checksum is used as an ID of URL.
func (l *Location) SHA256() []byte {

	s := sha256.Sum256([]byte(l.String()))
	return s[:]
}

// MarshalXML implements the xml.Marshaler interface.
//
// Example:
//
//	<loc>https://example.com/sample1.html</loc>
func (l *Location) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "loc"

	return e.EncodeElement(l.String(), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
//
// Example:
//
//	<loc>https://example.com/sample1.html</loc>
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
