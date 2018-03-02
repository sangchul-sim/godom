package tests

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestAddAttribute(t *testing.T) {
	var classTests = []struct {
		HTML     string
		attr     *html.Attribute
		selector string
		result   string
	}{
		{
			`<ul class="t1" id="item">`,
			&html.Attribute{
				"",
				"class",
				"t2",
			},
			"ul",
			`t1 t2`,
		},
		{
			`<p class="t1 t2">`,
			&html.Attribute{
				"",
				"title",
				"added title",
			},
			"p",
			`added title`,
		},
		{
			`<p class="t1 t2" title="">`,
			&html.Attribute{
				"",
				"title",
				"added title",
			},
			"p",
			`added title`,
		},
		{
			`<a data-label="Responsive Tab: All Products">All Products</a>`,
			&html.Attribute{
				"",
				"href",
				"https://developers.google.com/products/",
			},
			"a",
			`https://developers.google.com/products/`,
		},
		{
			`<a href="https://developers.google.com/products/">All Products</a>`,
			&html.Attribute{
				"",
				"href",
				"http://google.com",
			},
			"a",
			`https://developers.google.com/products/ http://google.com`,
		},
	}

	for idx, test := range classTests {
		s, err := cascadia.Compile(test.selector)
		if err != nil {
			t.Errorf("error compiling %q: %s", test.selector, err)
			continue
		}
		doc, err := html.Parse(strings.NewReader(test.HTML))
		if err != nil {
			t.Errorf("error parsing %q: %s", test.HTML, err)
			continue
		}

		qQuery := godom.NewGoQuery(s.MatchFirst(doc))
		qQuery.AddAttribute(test.attr)

		got := qQuery.GetAttributeByKey(test.attr.Key)
		if got.Val != test.result {
			t.Errorf("attrKey %s wanted %s, got %s instead at idx %d",
				test.attr.Key,
				test.result,
				got,
				idx,
			)
		}
	}
}
