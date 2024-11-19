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

	if *changefreq1 == "" {
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

	if *changefreq2 != *changefreq1 {
		t.Fatalf("FAIL: invalid changefreq: %#v\n", *changefreq2)
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
		t.Fatalf("FAIL: invalid prio2: %s\n", &prio2)
	}
}

func TestURL(t *testing.T) {

	url1 := sitemap.URL{
		Location:         sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency:  sitemap.ParseChangeFrequency("daily"),
		Priority:         sitemap.MustParsePriority("1.0"),
		LastModification: sitemap.MustParseLastModification("2006-01-02T15:04:05.999999999+07:00"),
		Images:           []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:       []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
		Comment:          sitemap.NewComment([]byte("comment1")),
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

	if *url2.Location != *url1.Location {
		t.Fatalf("Invalid Location: %s\n", url2.Location)
	}

	if *url2.ChangeFrequency != *url1.ChangeFrequency {
		t.Fatalf("Invalid ChangeFrequency: %s\n", *url2.ChangeFrequency)
	}

	if *url2.Priority != *url1.Priority {
		t.Fatalf("Invalid Priority: %s\n", url2.Priority)
	}

	if url2.LastModification.String() != url1.LastModification.String() {
		t.Fatalf("Invalid LastModification: %s\n", url2.LastModification)
	}

	if url2.Comment.String() != url1.Comment.String() {
		t.Fatalf("Invalid Comment: %s\n", url2.Comment)
	}
}

func TestURLSet(t *testing.T) {

	urls := sitemap.EmptyURLSet()

	urls.AppendURL(&sitemap.URL{
		Location:         sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency:  sitemap.ParseChangeFrequency("daily"),
		Priority:         sitemap.MustParsePriority("1.0"),
		LastModification: sitemap.MustParseLastModification("2006-01-02T15:04:05.999999999+07:00"),
		Images:           []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:       []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
	})

	urls.AppendURL(&sitemap.URL{
		Location:         sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency:  sitemap.ParseChangeFrequency("daily"),
		Priority:         sitemap.MustParsePriority("1.0"),
		LastModification: sitemap.MustParseLastModification("2006-01-02T15:04:05.999999999+07:00"),
		Images:           []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:       []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
	})

	t.Logf("\n%s\n\n", urls)
}
