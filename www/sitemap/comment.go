package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Comment xml.Comment

func NewComment(v string) Comment {
	if len(v) == 0 {
		return nil
	}

	return Comment(v)
}

func (c Comment) String() string {
	return string(c)
}

func (a Comment) Compare(b Comment) int {

	return bytes.Compare([]byte(a), []byte(b))
}

// MarshalXML implements the xml.Marshaler interface.
func (c Comment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	var v xml.Comment = xml.Comment(c)

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (c *Comment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v xml.Comment

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Comment: %w", err)
	}

	*c = Comment(v)

	return nil
}
