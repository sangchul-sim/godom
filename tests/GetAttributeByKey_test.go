package tests

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestGetAttributeByKey(t *testing.T) {
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
				"",
			},
			"ul",
			`t1`,
		},
		{
			`<p class="t1 t2">`,
			&html.Attribute{
				"",
				"class",
				"",
			},
			"p",
			`t1 t2`,
		},
		{
			`<a href="https://developers.google.com/products/">All Products</a>`,
			&html.Attribute{
				"",
				"href",
				"",
			},
			"a",
			`https://developers.google.com/products/`,
		},
		{
			`<a href="" data-label="Responsive Tab: All Products">All Products</a>`,
			&html.Attribute{
				"",
				"data-label",
				"",
			},
			"a",
			`Responsive Tab: All Products`,
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
