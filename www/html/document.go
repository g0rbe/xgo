package html

// type Document struct {
// 	head *goquery.Selection
// 	body *goquery.Selection
// }

// func ReadDocument(r io.Reader) (*Document, error) {

// 	// Load the HTML document
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	d := new(Document)

// 	d.head = doc.Find("head").Unwrap()

// 	d.body = doc.Find("body").Unwrap()

// 	return d, nil
// }

// func ParseDocument(b []byte) (*Document, error) {

// 	buf := bytes.NewReader(b)

// 	return ReadDocument(buf)
// }

// func (d *Document) HeadBytes() ([]byte, error) {

// 	v, err := d.head.Html()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return []byte(v), nil
// }

// func (d *Document) BodyBytes() ([]byte, error) {

// 	v, err := d.body.Html()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return []byte(v), nil
// }

// func (d *Document) BodyText() string {

// 	return d.body.Text()
// }
