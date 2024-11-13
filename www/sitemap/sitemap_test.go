package sitemap_test

import (
	"encoding/xml"
	"testing"

	"github.com/g0rbe/xgo/www/sitemap"
)

func TestAlternate(t *testing.T) {

	alt := sitemap.NewAlternate("hu", "https://example.com/hu/")

	v, err := xml.Marshal(alt)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	alt2 := new(sitemap.Alternate)

	err = xml.Unmarshal(v, alt2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if alt2.URL != alt.URL || alt2.Language != alt.Language {
		t.Fatalf("FAIL: invalid alternate: %#v\n", alt2)
	}
}

func TestImage(t *testing.T) {

	img, err := sitemap.ParseImageString("https://example.com/example.jpg")
	if err != nil {
		t.Fatalf("Failed to parse image string: %s\n", err)
	}

	v, err := xml.Marshal(img)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	img2 := new(sitemap.Image)

	err = xml.Unmarshal(v, img2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if img2.Location.String() != img.Location.String() {
		t.Fatalf("FAIL: invalid location: %#v\n", img2.Location.String())
	}
}

func TestLastModification(t *testing.T) {

	lastmod1, err := sitemap.ParseLastModification("2006-01-02T15:04:05.999999999+07:00")
	if err != nil {
		t.Fatalf("Failed to parse: %s\n", err)
	}

	v, err := xml.Marshal(lastmod1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	lastmod2 := new(sitemap.LastModification)

	err = xml.Unmarshal(v, lastmod2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if lastmod2.String() != lastmod1.String() {
		t.Fatalf("FAIL: invalid lastmod: %#v\n", *lastmod2)
	}
}

func TestLocation(t *testing.T) {

	url1, err := sitemap.ParseLocation("https://example.com/example/")
	if err != nil {
		t.Fatalf("Failed to parse: %s\n", err)
	}

	v, err := xml.Marshal(url1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	url2 := new(sitemap.Location)

	err = xml.Unmarshal(v, url2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if url2.String() != url1.String() {
		t.Fatalf("FAIL: invalid location: %#v\n", url2)
	}
}

func TestChangeFrequency(t *testing.T) {

	changefreq1 := sitemap.ParseChangeFrequency("always")

	if changefreq1 == "" {
		t.Fatalf("Failed to parse\n")
	}

	v, err := xml.Marshal(changefreq1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	changefreq2 := new(sitemap.ChangeFrequency)

	err = xml.Unmarshal(v, changefreq2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if *changefreq2 != changefreq1 {
		t.Fatalf("FAIL: invalid lastmod: %#v\n", *changefreq2)
	}
}

func TestPriority(t *testing.T) {

	prio1 := sitemap.Priority(0.5)

	v, err := xml.Marshal(prio1)
	if err != nil {
		t.Fatalf("Failed to marshal: %s\n", err)
	}

	t.Logf("\n%s\n", v)

	var prio2 sitemap.Priority

	err = xml.Unmarshal(v, &prio2)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %s\n", err)
	}

	if prio2 != prio1 {
		t.Fatalf("FAIL: invalid prio2: %s\n", prio2)
	}
}
