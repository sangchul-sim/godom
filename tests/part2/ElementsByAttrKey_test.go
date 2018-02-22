package part2

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestElementsByAttrKey(t *testing.T) {
	var classTests = []struct {
		HTML     string
		attr     *html.Attribute
		selector string
		results  []string
	}{
		{
			`<nav class="devsite-nav-responsive-tabs devsite-nav">
  <ul class="devsite-nav-list">
    <li class="devsite-nav-item devsite-nav-item-heading">
      <a href="https://developers.google.com/products/" data-label="Responsive Tab: All Products">
        All Products
      </a>
    </li>
    <li class="devsite-nav-item devsite-nav-item-heading">
      <a href="https://developers.google.com/products/develop/" data-label="Responsive Tab: Develop">
        Develop
      </a>
    </li>
  </ul>
</nav>`,
			&html.Attribute{"", "data-label", ""},
			`a[data-label!=""]`,
			[]string{
				`<a href="https://developers.google.com/products/" data-label="Responsive Tab: All Products">
        All Products
      </a>`,
				`<a href="https://developers.google.com/products/develop/" data-label="Responsive Tab: Develop">
        Develop
      </a>`,
			},
		},
		{
			`<ul class="t1 t2"><li class="t1" title="t1"></li><li class="t2" title="t2"></li></ul>`,
			&html.Attribute{"", "title", ""},
			`li[title!=""]`,
			[]string{
				`<li class="t1" title="t1"></li>`,
				`<li class="t2" title="t2"></li>`,
			},
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

		qQuery := godom.NewGoQuery(doc.FirstChild)
		matches := qQuery.ElementsByAttrKey(test.attr.Key)
		if len(matches) != len(test.results) {
			t.Errorf("attr %s wanted %d elements, got %d instead at idx %d",
				test.attr,
				len(test.results),
				len(matches),
				idx,
			)
			continue
		}
		for i, m := range matches {
			got := m.NodeString()
			if got != test.results[i] {
				t.Errorf("attr %s wanted %s, got %s instead at idx %d",
					test.selector,
					test.results,
					got,
					idx,
				)
			}
		}

		cascadiaMatches := s.MatchAll(doc)
		if len(matches) != len(test.results) {
			t.Errorf("selector %s wanted %d elements, got %d instead",
				test.selector,
				len(test.results),
				len(matches),
			)
			continue
		}
		for i, m := range cascadiaMatches {
			got := godom.NewGoQuery(m).NodeString()
			if got != test.results[i] {
				t.Errorf("selector %s wanted %s, got %s instead at idx %d",
					test.selector,
					test.results[i],
					got,
					idx,
				)
			}
		}
	}
}

// TODO val : Responsive ok, but val : Responsive Tab: ????
func TestElementsByAttrKeyVal(t *testing.T) {
	var classTests = []struct {
		HTML     string
		attr     *html.Attribute
		selector string
		results  []string
	}{
		{
			`<nav class="devsite-nav-responsive-tabs devsite-nav">
  <ul class="devsite-nav-list">
    <li class="devsite-nav-item devsite-nav-item-heading">
      <a href="https://developers.google.com/products/" data-label="Responsive Tab: All Products">
        All Products
      </a>
    </li>
    <li class="devsite-nav-item devsite-nav-item-heading">
      <a href="https://developers.google.com/products/develop/" data-label="Responsive Tab: Develop">
        Develop
      </a>
    </li>
  </ul>
</nav>`,
			&html.Attribute{"", "data-label", "Responsive"},
			`a[data-label="Responsive"]`,
			[]string{
				`<a href="https://developers.google.com/products/" data-label="Responsive Tab: All Products">
        All Products
      </a>`,
				`<a href="https://developers.google.com/products/develop/" data-label="Responsive Tab: Develop">
        Develop
      </a>`,
			},
		},
		{
			`<ul class="t1 t2"><li class="t1" title="t1"></li><li class="t2" title="t2"></li></ul>`,
			&html.Attribute{"", "title", "t2"},
			`li[title="t2"]`,
			[]string{
				`<li class="t2" title="t2"></li>`,
			},
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

		qQuery := godom.NewGoQuery(doc.FirstChild)
		matches := qQuery.ElementsByAttrKeyVal(test.attr.Key, test.attr.Val)
		if len(matches) != len(test.results) {
			t.Errorf("attr %s wanted %d elements, got %d instead at idx %d",
				test.attr,
				len(test.results),
				len(matches),
				idx,
			)
			continue
		}
		for i, m := range matches {
			got := m.NodeString()
			if got != test.results[i] {
				t.Errorf("attr %s wanted %s, got %s instead at idx %d",
					test.selector,
					test.results,
					got,
					idx,
				)
			}
		}

		sMatches := s.MatchAll(doc)
		if len(matches) != len(test.results) {
			t.Errorf("selector %s wanted %d elements, got %d instead",
				test.selector,
				len(test.results),
				len(matches),
			)
			continue
		}
		for i, m := range sMatches {
			got := godom.NewGoQuery(m).NodeString()
			if got != test.results[i] {
				t.Errorf("selector %s wanted %s, got %s instead at idx %d",
					test.selector,
					test.results[i],
					got,
					idx,
				)
			}
		}
	}
}
