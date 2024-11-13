package meta_test

import (
	"strings"
	"testing"

	"github.com/g0rbe/xgo/www/html/meta"
)

var TestDocument = `
<head>
	<meta name="generator" content="Hugo 0.135.0">
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width,initial-scale=1">
	<meta name="theme-color" content="rgb(38,38,38)">
	<meta name="robots" content="all">

	<title>Test Title</title>
	<meta name="title" content="Meta Test Title">

	<meta name="description" content="Test Description">

	<meta name="keywords" content="one two">

	<meta name="rating" content="adult">

	<link rel="canonical" href="https://example.com/">

	<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
	
	<link rel="alternate" hreflang="x-default" href="https://example.com/en/" />
    <link rel="alternate" hreflang="en" href="https://example.com/en/" />
    <link rel="alternate" hreflang="hu" href="https://example.com/hu/" />
</head>
	`

func TestMetaTitle(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if m.Title() != "Test Title" {
		t.Fatalf("FAIL: invalid title: %s\n", m.Title())
	}
}

func TestMetaDescription(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if m.Description() != "Test Description" {
		t.Fatalf("FAIL: invalid description: %s\n", m.Description())
	}
}

func TestMetaIcons(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Icons()) != 2 {
		t.Fatalf("FAIL: invalid icons: %#v\n", m.Icons())
	}
}

func TestMetaKeywords(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if m.Keywords() != "one two" {
		t.Fatalf("FAIL: invalid keywords: %s\n", m.Keywords())
	}
}

func TestMetaRating(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Rating()) != 1 && m.Rating()[0] != "adult" {
		t.Fatalf("FAIL: invalid rating: %s\n", m.Keywords())
	}
}

func TestMetaCanonical(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if m.Canonical() != "https://example.com/" {
		t.Fatalf("FAIL: Invalid canonical: %s\n", m.Canonical())
	}
}

func TestMetaAlternate(t *testing.T) {

	m, err := meta.ReadDocument(strings.NewReader(TestDocument))
	if err != nil {
		t.Fatalf("ERROR: %s\n", err)
	}

	if len(m.Alternate()) != 3 {
		t.Fatalf("FAIL: empty alternates")
	}

	t.Logf("%#v\n", m.Alternate())
}
