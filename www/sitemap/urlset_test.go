package sitemap_test

import (
	"testing"

	"github.com/g0rbe/xgo/www/sitemap"
)

func TestURLSet(t *testing.T) {

	urls := sitemap.EmptyURLSet()

	urls.AppendURL(&sitemap.URL{
		Location:        sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency: sitemap.ParseChangeFrequency("daily"),
		Priority:        sitemap.Priority(1.0),
		Images:          []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:      []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
	})

	urls.AppendURL(&sitemap.URL{
		Location:        sitemap.MustParseLocation("https://example.com/"),
		ChangeFrequency: sitemap.ParseChangeFrequency("daily"),
		Priority:        sitemap.Priority(1.0),
		Images:          []sitemap.Image{sitemap.MustParseImageString("https://example.com/cover.jpg")},
		Alternates:      []sitemap.Alternate{sitemap.NewAlternate("https://example.hu/", "hu")},
	})

	t.Logf("\n%s\n\n", urls)
}
