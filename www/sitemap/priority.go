package sitemap

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// Priority type used for the priority field
//
//	<priority>1.0</priority>
type Priority float64

func (p Priority) String() string {
	return strconv.FormatFloat(float64(p), 'f', 1, 64)
}

func (p Priority) IsEmpty() bool {
	return p == 0
}

func (p Priority) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "priority"

	return e.EncodeElement(p.String(), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (p *Priority) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return fmt.Errorf("failed to decode priority: %w", err)
	}

	v = strings.TrimSpace(v)

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return fmt.Errorf("failed to parse priority \"%s\": %w", v, err)

	}

	if f < 0.0 || f > 1.0 {
		return fmt.Errorf("invalid value for priority: %f", f)
	}

	*p = (Priority)(f)

	return nil
}
