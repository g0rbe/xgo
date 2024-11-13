package sitemap_test

import (
	"encoding/xml"
	"testing"

	"github.com/g0rbe/xgo/www/sitemap"
)

func TestURL(t *testing.T) {

	url1 := sitemap.URL{
		Location:        sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency: sitemap.ParseChangeFrequency("daily"),
		Priority:        sitemap.Priority(1.0),
		Images:          []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:      []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
		Comment:         "comment1",
	}

	v1, err := xml.MarshalIndent(url1, "    ", "    ")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("\n%s\n", v1)

	url2 := new(sitemap.URL)

	err = xml.Unmarshal(v1, url2)
	if err != nil {
		t.Fatalf("Unmarshal error: %s\n", err)
	}

	v2, err := xml.MarshalIndent(url2, "    ", "    ")
	if err != nil {
		t.Fatalf("FAIL: %s\n", err)
	}

	t.Logf("\n%s\n", v2)

	if url2.Location.String() != url1.Location.String() {
		t.Fatalf("Invalid Location: %s\n", url2.Location)
	}

	if url2.ChangeFrequency != url1.ChangeFrequency {
		t.Fatalf("Invalid ChangeFrequency: %s\n", url2.ChangeFrequency)
	}

	if url2.Priority != url1.Priority {
		t.Fatalf("Invalid Priority: %s\n", url2.Priority)
	}

	if url2.Comment != url1.Comment {
		t.Fatalf("Invalid Comment: %s\n", url2.Comment)
	}
}
