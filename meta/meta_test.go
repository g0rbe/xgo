package meta_test

import (
	"testing"

	"github.com/g0rbe/xgo/meta"
)

func TestMetaTitle(t *testing.T) {

	m, err := meta.FetchHeader("https://gorbe.io/about")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Title()) == 0 {
		t.Fatalf("FAIL: empty title\n")
	}
}

func TestMetaDescription(t *testing.T) {

	m, err := meta.FetchHeader("https://gorbe.io/about")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Description()) == 0 {
		t.Fatalf("FAIL: empty description\n")
	}
}

func TestMetaIcons(t *testing.T) {

	m, err := meta.FetchHeader("https://gorbe.io/about")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Icons()) == 0 {
		t.Fatalf("FAIL: empty icons")
	}
}

func TestMetaKeywords(t *testing.T) {

	m, err := meta.FetchHeader("https://www.w3schools.com/tags/tag_meta.asp")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Keywords()) == 0 {
		t.Fatalf("FAIL: Empty keywords: %#v\n", m.Keywords())
	}
}

func TestMetaRating(t *testing.T) {

	m, err := meta.FetchHeader("https://www.pornhub.com/")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Rating()) == 0 {
		t.Fatalf("FAIL: Empty rating: %#v\n", m.Rating())
	}
}

func TestMetaCanonical(t *testing.T) {

	m, err := meta.FetchHeader("https://gorbe.io")
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if m.Canonical() != "https://gorbe.io/" {
		t.Fatalf("FAIL: Invalid canonical: %s\n", m.Canonical())
	}
}
