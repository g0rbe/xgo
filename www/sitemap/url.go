package sitemap

import (
	"encoding/xml"
	"net/url"
)

// URL is the <url> field of the sitemap
//
//		<url>
//	    	<loc>https://example.com/</loc>
//	    	<lastmod>0001-01-01T00:00:00Z</lastmod>
//	    	<changefreq>daily</changefreq>
//			<priority>1.0</priority>
//			<image:image>
//				<image:loc>https://example.com/cover.jpg</image:loc>
//			</image:image>
//			<xhtml:link rel="alternate" hreflang="https://example.hu/" href="hu"></xhtml:link>
//			<!--comment1-->
//		</url>
type URL struct {
	id               []byte
	Location         *Location
	LastModification *LastModification
	ChangeFrequency  *ChangeFrequency
	Priority         *Priority
	Images           []Image
	Alternates       []Alternate
	Comment          Comment
}

func EmptyURL() *URL {
	return new(URL)
}

func (u *URL) AddLocationURL(l *url.URL) {

	if l == nil {
		return
	}

	u.Location = (*Location)(l)
	u.id = u.Location.SHA256()
}

func (u *URL) AppendAlternate(href, hreflang string) {
	if href == "" || hreflang == "" {
		return
	}

	u.Alternates = append(u.Alternates, NewAlternate(href, hreflang))
}

// MarshalXML implements the xml.Marshaler interface.
func (u URL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "url"

	v := struct {
		XMLName          xml.Name          `xml:"url"`
		Location         *Location         `xml:"loc"`
		LastModification *LastModification `xml:"lastmod,omitempty"`
		ChangeFrequency  *ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority         *Priority         `xml:"priority,omitempty"`
		Images           []Image           `xml:"image:image,omitempty"`
		Alternates       []Alternate       `xml:"xhtml:link,omitempty"`
		Comment          xml.Comment       `xml:",comment"`
	}{
		Location:         u.Location,
		LastModification: u.LastModification,
		ChangeFrequency:  u.ChangeFrequency,
		Priority:         u.Priority,
		Images:           u.Images,
		Alternates:       u.Alternates,
		Comment:          xml.Comment(u.Comment),
	}

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		XMLName          xml.Name          `xml:"url"`
		Location         *Location         `xml:"loc"`
		LastModification *LastModification `xml:"lastmod,omitempty"`
		ChangeFrequency  *ChangeFrequency  `xml:"changefreq,omitempty"`
		Priority         *Priority         `xml:"priority,omitempty"`
		Images           []Image           `xml:"image,omitempty"`
		Alternates       []Alternate       `xml:"link,omitempty"`
		Comment          xml.Comment       `xml:",comment"`
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	u.id = v.Location.SHA256()

	u.Location = v.Location
	u.LastModification = v.LastModification
	u.ChangeFrequency = v.ChangeFrequency
	u.Priority = v.Priority
	u.Images = v.Images
	u.Alternates = v.Alternates
	u.Comment = Comment(v.Comment)

	return nil
}
