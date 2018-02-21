package tests

import (
	"strings"
	"testing"

	"github.com/andybalholm/cascadia"
	"github.com/sangchul-sim/godom"
	"golang.org/x/net/html"
)

func TestGetElementByID(t *testing.T) {
	var classTests = []struct {
		HTML     string
		id       string
		selector string
		result   string
	}{
		{
			`<ul><li class="t1" id="item"><li class="t2">`,
			"item",
			"#item",
			`<li class="t1" id="item"></li>`,
		},
		{
			`<p class="t1 t2">`,
			"item",
			"#item",
			`not found`,
		},
		{
			`<ul class="top">
				<li class="item" id="active">item1</li>
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
				<li class="top item" id="active">item4</li>
				<li>item5</li>
			</ul>`,
			"active",
			"#active",
			`<li class="item" id="active">item1</li>`,
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

		qQuery := godom.NewGoQuery(doc)
		matche, err := qQuery.GetElementByID(test.id)
		if err != nil {
			if err.Error() != test.result {
				t.Errorf("test.result %s, err %s", test.result, err)
			}
			continue
		}
		got := godom.NewGoQuery(matche).NodeString()
		if got != test.result {
			t.Errorf("class %s wanted %s, got %s instead at idx %d",
				test.id,
				test.result[idx],
				got,
				idx,
			)
		}

		sMatche := s.MatchFirst(doc)
		got = godom.NewGoQuery(sMatche).NodeString()
		if got != test.result {
			t.Errorf("selector %s wanted %s, got %s instead at idx %d",
				test.id,
				test.result,
				got,
				idx,
			)
		}
	}
}
