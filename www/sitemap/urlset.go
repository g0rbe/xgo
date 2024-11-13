package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"sync"
)

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

// MarshalXML implements the xml.Marshaler interface.
func (s URLSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "urlset"

	v := struct {
		SitemapNS string `xml:"xmlns,attr"`
		XHTMLAttr string `xml:"xmlns:xhtml,attr"`
		ImageAttr string `xml:"xmlns:image,attr"`
		URLs      []URL  `xml:"url"`
	}{
		SitemapNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		XHTMLAttr: XHTMLNS,
		ImageAttr: ImageNS,
		URLs:      s.URLs,
	}

	return e.EncodeElement(v, start)
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (s *URLSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	//start.Name.Local = "xhtml:link"

	v := struct {
		URLs []URL `xml:"url,omitempty"`
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	// if u == nil {
	// 	u = new(URL)
	// }

	s.URLs = v.URLs

	return nil
}
