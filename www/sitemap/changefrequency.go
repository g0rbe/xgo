package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// ChangeFrequency is the changefreq field of the sitemap
//
//	<changefreq>always</changefreq>
type ChangeFrequency string

// Valid values for ChangeFrequency
var (
	ChangeFreqAlways  ChangeFrequency = "always"
	ChangeFreqHourly  ChangeFrequency = "hourly"
	ChangeFreqDaily   ChangeFrequency = "daily"
	ChangeFreqWeekly  ChangeFrequency = "weekly"
	ChangeFreqMonthly ChangeFrequency = "monthly"
	ChangeFreqYearly  ChangeFrequency = "yearly"
	ChangeFreqNever   ChangeFrequency = "never"
)

// ParseChangeFrequency parses ChangeFrequency from v.
//
// If v is not a valid ChangeFrequency, returns nil
func ParseChangeFrequency(v string) *ChangeFrequency {

	switch v {
	case "always":
		return &ChangeFreqAlways
	case "hourly":
		return &ChangeFreqHourly
	case "daily":
		return &ChangeFreqDaily
	case "weekly":
		return &ChangeFreqWeekly
	case "monthly":
		return &ChangeFreqMonthly
	case "yearly":
		return &ChangeFreqYearly
	case "never":
		return &ChangeFreqNever
	default:
		return nil
	}
}

func (f *ChangeFrequency) String() string {
	return string(*f)
}

func (f *ChangeFrequency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "changefreq"

	return e.EncodeElement(f.String(), start)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (f *ChangeFrequency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return fmt.Errorf("failed to decode ChangeFrequency: %w", err)
	}

	if v == "" {
		return fmt.Errorf("empty value for changefreq")
	}

	v = strings.TrimSpace(v)

	c := ParseChangeFrequency(v)

	if c == nil {
		return fmt.Errorf("invalid value for changefreq: %s", c)
	}

	*f = *c

	return nil
}
