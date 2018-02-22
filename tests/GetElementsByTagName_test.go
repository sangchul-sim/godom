package tests

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestGetElementsByTagName(t *testing.T) {
	var tagTests = []struct {
		HTML     string
		tagName  string
		selector string
		results  []string
	}{
		{
			`<p class="t1 t2">`,
			"p",
			"p",
			[]string{
				`<p class="t1 t2"></p>`,
			},
		},
		{
			`<ul class="top">
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
			</ul>`,
			"li",
			"li",
			[]string{
				`<li class="item" val="item1">item1</li>`,
				`<li class="item" val="item2">item2</li>`,
				`<li val="item3">
					<ul>
						<li class="item top" val="item3-1">item3-1</li>
						<li val="item3-2">
							<ul>
								<li class="item" val="item3-2-1">item3-2-1</li>
								<li class="top" val="item3-2-2">item3-2-2</li>
							</ul>
						</li>
					</ul>
				</li>`,
				`<li class="item top" val="item3-1">item3-1</li>`,
				`<li val="item3-2">
							<ul>
								<li class="item" val="item3-2-1">item3-2-1</li>
								<li class="top" val="item3-2-2">item3-2-2</li>
							</ul>
						</li>`,
				`<li class="item" val="item3-2-1">item3-2-1</li>`,
				`<li class="top" val="item3-2-2">item3-2-2</li>`,
				`<li class="top item" val="item4">item4</li>`,
				`<li val="item5">item5</li>`,
			},
		},
	}

	for idx, test := range tagTests {
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

		qQuery := godom.NewGoQuery(doc)
		matches := qQuery.GetElementsByTagName(test.tagName)
		if len(matches) != len(test.results) {
			t.Errorf("class %s wanted %d elements, got %d instead at idx %d",
				test.tagName,
				len(test.results),
				len(matches),
				idx,
			)
			continue
		}
		for i, m := range matches {
			got := godom.NewGoQuery(m).NodeString()
			if got != test.results[i] {
				t.Errorf("class %s wanted %s, got %s instead at idx %d",
					test.tagName,
					test.results[i],
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
