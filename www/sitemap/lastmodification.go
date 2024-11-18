package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

type LastModification time.Time

func ParseLastModification(s string) (*LastModification, error) {

	var layout string

	switch len(s) {
	case 10:
		layout = time.DateOnly
	case 20:
		layout = "2006-01-02T15:04:05Z"
	case 24:
		layout = "2006-01-02T15:04:05.999Z"
	case 25:
		layout = "2006-01-02T15:04:05.999999999Z07:00"
	default:
		layout = time.RFC3339Nano
	}

	v, err := time.Parse(layout, s)
	if err != nil {
		return nil, err
	}

	return (*LastModification)(&v), nil
}

func MustParseLastModification(s string) *LastModification {

	v, err := ParseLastModification(s)
	if err != nil {
		panic(err)
	}

	return v
}

func (t *LastModification) Format(layout string) string {
	return time.Time(*t).Format(layout)
}

// String returns the RFC3339Nano formatted time of t.
func (t *LastModification) String() string {
	return t.Format(time.RFC3339Nano)
}

func (t *LastModification) IsZero() bool {
	return time.Time(*t).IsZero()
}

func (l *LastModification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if start.Name.Local != "lastmod" {
		start.Name.Local = "lastmod"
	}

	return e.EncodeElement(l.String(), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (t *LastModification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var s string

	err := d.DecodeElement(&s, &start)
	if err != nil {
		return fmt.Errorf("failed to decode LastModification: %w", err)
	}

	v, err := ParseLastModification(strings.TrimSpace(s))
	if err != nil {
		return err
	}

	*t = *v

	return err
}
