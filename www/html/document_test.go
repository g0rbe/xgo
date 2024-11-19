package html_test

// var TestDocument = `<!DOCTYPE html>
// <html>
// 	<head>
// 		<title>G⌬RBE</title>
// 	</head>
// 	<body>
// 		<p>Test Body</p>
// 	</body>
// </html>
// `

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
