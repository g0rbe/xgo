package sitemap

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

// ImageNS is the image namespace.
const ImageNS = "http://www.google.com/schemas/sitemap-image/1.1"

type Image struct {
	Location *url.URL
}

func ParseImageString(loc string) (Image, error) {

	imageLoc, err := url.Parse(loc)
	if err != nil {
		return Image{}, err
	}
	return Image{Location: imageLoc}, nil
}

func MustParseImageString(loc string) Image {

	imageLoc, err := url.Parse(loc)
	if err != nil {
		panic(err)
	}

	return Image{Location: imageLoc}
}

func (i Image) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	// Set tag
	start.Name.Local = "image:image"

	v := struct {
		Location string `xml:"image:loc"`
	}{
		Location: i.Location.String(),
	}

	return e.EncodeElement(v, start)
}

func (i *Image) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := struct {
		XMLName  xml.Name `xml:"image image"`
		Location struct {
			XMLName xml.Name `xml:"image loc"`
			Loc     string   `xml:",chardata"`
		}
	}{}

	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	imageLoc, err := url.Parse(v.Location.Loc)
	if err != nil {
		return fmt.Errorf("failed to parse url: %w", err)
	}

	i.Location = imageLoc

	return nil
}
