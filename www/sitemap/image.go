package sitemap

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

// ImageNS is the image namespace.
const ImageNS = "http://www.google.com/schemas/sitemap-image/1.1"

// Image is the image extension of the sitemap.
//
//	<image:image>
//	  <image:loc>https://example.com/picture.jpg</image:loc>
//	</image:image>
//
// See: https://developers.google.com/search/docs/crawling-indexing/sitemaps/image-sitemaps
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

func (i Image) String() string {
	return i.Location.String()
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

	if start.Name.Space == ImageNS {
		start.Name.Space = "image"
	}

	v := struct {
		XMLName  xml.Name `xml:"image image"`
		Location struct {
			XMLName xml.Name `xml:"loc"`
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
