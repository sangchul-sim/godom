package tests

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func aTestParse(t *testing.T) {
	text := `<ul class="top">
				<li class="item" val="item1">item1</li>
				<li class="item" val="item2">item2</li>
				<li val="item3">
					<ul>
						<li class="item top" val="item3-1">item3-1</li>
						<li val="item3-2">
							<ul>
								<li class="item" val="item3-2-1">item3-2-1</li>
								<li class="top" val="item3-2-2">item3-2-2</li>
							</ul>
						</li>
					</ul>
				</li>
				<li class="top item" val="item4">item4</li>
				<li val="item5">item5</li>
			</ul>`

	//doc, err := html.Parse(strings.NewReader(text))
	doc, err := html.Parse(bytes.NewReader([]byte(text)))
	if err != nil {
		panic(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		fmt.Println("c.Data", n.Data)
		/**
		ErrorNode NodeType = iota
		TextNode
		DocumentNode
		ElementNode
		CommentNode
		DoctypeNode
		scopeMarkerNode
		*/
		switch n.Type {
		case 0:
			fmt.Println("c.Val", "ErrorNode")
		case 1:
			fmt.Println("c.Val", "TextNode")
		case 2:
			fmt.Println("c.Val", "DocumentNode")
		case 3:
			fmt.Println("c.Val", "ElementNode")
		case 4:
			fmt.Println("c.Val", "CommentNode")
		case 5:
			fmt.Println("c.Val", "DoctypeNode")
		case 6:
			fmt.Println("c.Val", "scopeMarkerNode")
		}

		fmt.Println("c.Attr", n.Attr)
		fmt.Println("\n")

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	//for c := doc.FirstChild; c != nil; c = c.NextSibling {
	//	fmt.Println("c.Data", c.Data)
	//	fmt.Println("c.Attr", c.Attr)
	//}

	//z := html.NewTokenizer(bytes.NewReader([]byte(text)))
	//for {
	//	tt := z.Next()
	//	fmt.Println("tt.String()", tt.String())
	//
	//	fmt.Println("z.Text()", string(z.Text()))
	//	name, hasAttr := z.TagName()
	//	fmt.Printf("z.TagName() %s %v\n", string(name), hasAttr)
	//
	//	if tt == html.ErrorToken {
	//		break
	//	}
	//}
}

func TestGetElementsByAttrKey(t *testing.T) {
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

		qQuery := godom.NewGoDom(doc.FirstChild)
		matches := qQuery.GetElementsByAttrKey(test.attr.Key)
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
			got := godom.NewGoDom(m).NodeString()
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
			got := godom.NewGoDom(m).NodeString()
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
func TestGetElementsByAttrKeyVal(t *testing.T) {
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

		qQuery := godom.NewGoDom(doc.FirstChild)
		matches := qQuery.GetElementsByAttrKeyVal(test.attr.Key, test.attr.Val)
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
			got := godom.NewGoDom(m).NodeString()
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
			got := godom.NewGoDom(m).NodeString()
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
