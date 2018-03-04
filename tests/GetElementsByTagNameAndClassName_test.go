package tests

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestGetElementsByTagNameAndClassName(t *testing.T) {
	var classTests = []struct {
		HTML     string
		tag      string
		attr     *html.Attribute
		selector string
		results  []string
	}{
		{
			`<ul><li class="t1"><li class="t2">`,
			"li",
			&html.Attribute{"", godom.AttrKeyClass, "t1"},
			"li.t1",
			[]string{
				`<li class="t1"></li>`,
			},
		},
		{
			`<p class="t1 t2">`,
			"p",
			&html.Attribute{"", godom.AttrKeyClass, "t1"},
			"p.t1",
			[]string{
				`<p class="t1 t2"></p>`,
			},
		},
		{
			`<div class="test">`,
			"div",
			&html.Attribute{"", godom.AttrKeyClass, "teST"},
			"div.teST",
			[]string{},
		},
		{
			`<p class="t1 t2">`,
			"p",
			&html.Attribute{"", godom.AttrKeyClass, "t1.fail"},
			"p.t1.fail",
			[]string{},
		},
		{
			`<ul class="top">
				<li class="item">item1</li>
				<li class="item">item2</li>
				<li>
					<ul>
						<li class="item top">item3-1</li>
						<li>
							<ul class="item">
								<li class="item">item3-2-1</li>
								<li class="top">item3-2-2</li>
							</ul>
						</li>
					</ul>
				</li>
				<li class="top item">item4</li>
				<li>item5</li>
			</ul>`,
			"li",
			&html.Attribute{"", godom.AttrKeyClass, "item"},
			"li.item",
			[]string{
				`<li class="item">item1</li>`,
				`<li class="item">item2</li>`,
				`<li class="item top">item3-1</li>`,
				`<li class="item">item3-2-1</li>`,
				`<li class="top item">item4</li>`,
			},
		},
		{
			`<p class="">This text should be green.</p><p>This text should be green.</p>`,
			"p",
			&html.Attribute{"", godom.AttrKeyClass, ""},
			`p[class=""]`,
			[]string{
				`<p class="">This text should be green.</p>`,
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

		qQuery := godom.NewGoDom(doc)
		matches := qQuery.GetElementsByTagNameAndClassName(test.tag, test.attr.Val)
		if len(matches) != len(test.results) {
			t.Errorf("class %s wanted %d elements, got %d instead at idx %d",
				test.attr.Val,
				len(test.results),
				len(matches),
				idx,
			)
			continue
		}
		for i, m := range matches {
			got := godom.NewGoDom(m).NodeString()
			if got != test.results[i] {
				t.Errorf("class %s wanted %s, got %s instead at idx %d",
					test.attr.Val,
					test.results[i],
					got,
					idx,
				)
			}
		}

		sMatches := s.MatchAll(doc)
		if len(matches) != len(test.results) {
			t.Errorf("class %s wanted %d elements, got %d instead",
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
