package html_test

import (
	"testing"

	"github.com/g0rbe/xgo/www/html"
)

var TestDocument = `<!DOCTYPE html>
<html>
	<head>
		<title>G⌬RBE</title>
	</head>
	<body>
		<h1>Test Body</h1>
		<h2>Test Sub Body</h2>
		<p class="removeme">Test Class</p>
		<p class="deleteme">Test Class</p>
	</body>
</html>
`

// func TestDocumentParse(t *testing.T) {

// 	doc, err := html.ParseDocument([]byte(TestDocument))
// 	if err != nil {
// 		t.Fatalf("Failed to parse document: %s\n", err)
// 	}

// 	headBytes, err := doc.HeadBytes()
// 	if err != nil {
// 		t.Fatalf("Failed to get head bytes: %s\n", err)
// 	}

// 	t.Logf("%s\n", headBytes)

// 	bodyBytes, err := doc.BodyBytes()
// 	if err != nil {
// 		t.Fatalf("Failed to get body bytes: %s\n", err)
// 	}

// 	t.Logf("%s\n", bodyBytes)
// }

func TestBodyRemove(t *testing.T) {

	b, err := html.ParseBody([]byte(TestDocument))
	if err != nil {
		t.Fatalf("Failed to parse body: %s\n", err)
	}

	bodyHtml, err := b.HTML()
	if err != nil {
		t.Fatalf("Failed to get body HTML: %s\n", err)
	}

	t.Logf("\n%s\n", bodyHtml)

	b.RemoveSelections([]string{".removeme", ".deleteme"})

	bodyHtml, err = b.HTML()
	if err != nil {
		t.Fatalf("Failed to get body HTML: %s\n", err)
	}

	t.Logf("\n%s\n", bodyHtml)
}
