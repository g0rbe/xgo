package sitemap

import (
	"encoding/xml"
)

type URL struct {
	Location         *Location
	LastModification LastModification
	ChangeFrequency  ChangeFrequency
	Priority         Priority
	Images           []Image
	Alternates       []Alternate
	Comment          string
}

func EmptyURL() *URL {
	return new(URL)
}

// func (u *URL) AppendImage(loc string) {
// 	if loc == "" {
// 		return
// 	}

// 	u.Images = append(u.Images, NewImage(loc))
// }

func (u *URL) AppendAlternate(href, hreflang string) {
	if href == "" || hreflang == "" {
		return
	}

	u.Alternates = append(u.Alternates, NewAlternate(href, hreflang))
}

// MarshalXML implements the xml.Marshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (u URL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "url"

	v := struct {
		Location         string      `xml:"loc"`
		LastModification string      `xml:"lastmod,omitempty"`
		ChangeFrequency  string      `xml:"changefreq,omitempty"`
		Priority         string      `xml:"priority,omitempty"`
		Images           []Image     `xml:"image:image,omitempty"`
		Alternates       []Alternate `xml:"xhtml:link,omitempty"`
		Comment          xml.Comment `xml:",comment"`
	}{
		Location:         u.Location.String(),
		LastModification: u.LastModification.String(),
		ChangeFrequency:  u.ChangeFrequency.String(),
		Priority:         u.Priority.String(),
		Images:           u.Images,
		Alternates:       u.Alternates,
		Comment:          xml.Comment(u.Comment),
	}

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
//
// Example:
//
//	<xhtml:link rel="alternate" hreflang="hu" href="https://example.com/hu/"></xhtml:link>
func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		XMLName          xml.Name         `xml:"url"`
		Location         *Location        `xml:"loc"`
		LastModification LastModification `xml:"lastmod,omitempty"`
		ChangeFrequency  ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority         Priority         `xml:"priority,omitempty"`
		Images           []Image          `xml:"image image,omitempty"`
		Alternates       []Alternate      `xml:"xhtml link,omitempty"`
		Comment          xml.Comment      `xml:",comment"`
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	// if u == nil {
	// 	u = new(URL)
	// }

	u.Location = v.Location
	u.LastModification = v.LastModification
	u.ChangeFrequency = v.ChangeFrequency
	u.Priority = v.Priority
	u.Images = v.Images
	u.Alternates = v.Alternates
	u.Comment = string(v.Comment)

	return nil
}
