package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sync"
)

const SitemapNS = "http://www.sitemaps.org/schemas/sitemap/0.9"

type URLSet struct {
	URLs []URL
	m    *sync.RWMutex
}

func EmptyURLSet() *URLSet {
	return &URLSet{m: new(sync.RWMutex)}
}

func ReadURLSet(r io.Reader) (*URLSet, error) {

	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := EmptyURLSet()

	err = xml.Unmarshal(buf, s)

	return s, err
}

func (s *URLSet) AppendURL(u *URL) {

	s.m.Lock()
	defer s.m.Unlock()

	for i := range s.URLs {
		if s.URLs[i].Location == u.Location {
			return
		}
	}

	s.URLs = append(s.URLs, *u)

}

func (s *URLSet) ToXML() ([]byte, error) {

	s.m.RLock()
	defer s.m.RUnlock()

	v, err := xml.Marshal(s)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.WriteString(xml.Header)
	buf.Write(v)

	return buf.Bytes(), nil
}

func (s *URLSet) String() string {

	s.m.RLock()
	defer s.m.RUnlock()

	v, err := xml.MarshalIndent(s, "", "    ")
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}

	return xml.Header + string(v)
}

// HasAlternate returns true if any URL has at least one Alternate field set.
func (s *URLSet) HasAlternate() bool {

	s.m.RLock()
	defer s.m.RUnlock()

	for i := range s.URLs {
		if len(s.URLs[i].Alternates) > 0 {
			return true
		}
	}

	return false
}

// HasImage returns true if any URL has at least one Image field set.
func (s *URLSet) HasImage() bool {

	s.m.RLock()
	defer s.m.RUnlock()

	for i := range s.URLs {
		if len(s.URLs[i].Images) > 0 {
			return true
		}
	}

	return false
}

// MarshalXML implements the xml.Marshaler interface.
func (s URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	s.m.RLock()
	defer s.m.RUnlock()

	if start.Name.Local != "urlset" {
		start.Name.Local = "urlset"
	}

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: SitemapNS})

	// Append XHTML namespace only if any alternate exists
	if s.HasAlternate() {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xhtml"}, Value: XHTMLNS})
	}

	// Append image namespace only if any mage exists
	if s.HasImage() {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:image"}, Value: ImageNS})
	}

	v := struct {
		URLs []URL `xml:"url,omitempty"`
	}{
		URLs: s.URLs,
	}

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (s *URLSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		URLs []URL `xml:"url,omitempty"`
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	s.URLs = v.URLs

	return nil
}
