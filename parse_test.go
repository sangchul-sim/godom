package dom

import (
	"bytes"
	"strings"
	"testing"

	"fmt"

	"golang.org/x/net/html"
)

func nodeString(n *html.Node) string {
	buf := bytes.NewBufferString("")
	html.Render(buf, n)
	return buf.String()
}

//func TestGetElementsByClassName(t *testing.T) {
//	var classTests = []struct {
//		HTML     string
//		class    string
//		selector string
//		results  []string
//	}{
//		{
//			`<ul><li class="t1"><li class="t2">`,
//			"t1",
//			".t1",
//			[]string{
//				`<li class="t1"></li>`,
//			},
//		},
//		{
//			`<p class="t1 t2">`,
//			"t1",
//			"p.t1",
//			[]string{
//				`<p class="t1 t2"></p>`,
//			},
//		},
//		{
//			`<div class="test">`,
//			"teST",
//			"div.teST",
//			[]string{},
//		},
//		{
//			`<p class="t1 t2">`,
//			"t1.fail",
//			".t1.fail",
//			[]string{},
//		},
//		{
//			`<ul class="top">
//				<li class="item">item1</li>
//				<li class="item">item2</li>
//				<li>
//					<ul>
//						<li class="item top">item3-1</li>
//						<li>
//							<ul class="item">
//								<li class="item">item3-2-1</li>
//								<li class="top">item3-2-2</li>
//							</ul>
//						</li>
//					</ul>
//				</li>
//				<li class="top item">item4</li>
//				<li>item5</li>
//			</ul>`,
//			"item",
//			".item",
//			[]string{
//				`<li class="item">item1</li>`,
//				`<li class="item">item2</li>`,
//				`<li class="item top">item3-1</li>`,
//				`<ul class="item">
//								<li class="item">item3-2-1</li>
//								<li class="top">item3-2-2</li>
//							</ul>`,
//				`<li class="item">item3-2-1</li>`,
//				`<li class="top item">item4</li>`,
//			},
//		},
//		{
//			`<p class="">This text should be green.</p><p>This text should be green.</p>`,
//			"",
//			`p[class=""]`,
//			[]string{
//				`<p class="">This text should be green.</p>`,
//			},
//		},
//	}
//
//	for idx, test := range classTests {
//		s, err := cascadia.Compile(test.selector)
//		if err != nil {
//			t.Errorf("error compiling %q: %s", test.selector, err)
//			continue
//		}
//		doc, err := html.Parse(strings.NewReader(test.HTML))
//		if err != nil {
//			t.Errorf("error parsing %q: %s", test.HTML, err)
//			continue
//		}
//
//		matches := GetElementsByClassName(doc, test.class)
//		if len(matches) != len(test.results) {
//			t.Errorf("class %s wanted %d elements, got %d instead at idx %d",
//				test.class,
//				len(test.results),
//				len(matches),
//				idx,
//			)
//			continue
//		}
//		for i, m := range matches {
//			got := nodeString(m)
//			if got != test.results[i] {
//				t.Errorf("class %s wanted %s, got %s instead at idx %d",
//					test.class,
//					test.results[i],
//					got,
//					idx,
//				)
//			}
//		}
//
//		cascadiaMatches := s.MatchAll(doc)
//		if len(matches) != len(test.results) {
//			t.Errorf("class %s wanted %d elements, got %d instead",
//				test.selector,
//				len(test.results),
//				len(matches),
//			)
//			continue
//		}
//		for i, m := range cascadiaMatches {
//			got := nodeString(m)
//			if got != test.results[i] {
//				t.Errorf("selector %s wanted %s, got %s instead at idx %d",
//					test.selector,
//					test.results[i],
//					got,
//					idx,
//				)
//			}
//		}
//	}
//}

func TestHasClassName(t *testing.T) {
	var classTests = []struct {
		HTML     string
		class    string
		selector string
		results  bool
	}{
		//{
		//	`<ul><li class="t1"><li class="t2">`,
		//	"t1",
		//	".t1",
		//	true,
		//},
		{
			`<ul class="t1 t2"><li class="t1"></li><li class="t2"></li></ul>`,
			"t1",
			"p.t1",
			true,
		},
		//{
		//	`<div class="test">`,
		//	"teST",
		//	"div.teST",
		//	false,
		//},
		//{
		//	`<p class="t1 t2">`,
		//	"t1.fail",
		//	".t1.fail",
		//	false,
		//},
		//{
		//	`<p class="">This text should be green.</p><p>This text should be green.</p>`,
		//	"",
		//	`p[class=""]`,
		//	true,
		//},
	}

	for idx, test := range classTests {
		//s, err := cascadia.Compile(test.selector)
		//if err != nil {
		//	t.Errorf("error compiling %q: %s", test.selector, err)
		//	continue
		//}
		fmt.Println("test.HTML", test.HTML)
		doc, err := html.Parse(strings.NewReader(test.HTML))
		if err != nil {
			t.Errorf("error parsing %q: %s", test.HTML, err)
			continue
		}
		fmt.Println(doc)

		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			fmt.Println("data", c.Data)
			fmt.Println("attri", c.Attr)
		}

		got := HasClassName(doc.FirstChild, test.class)
		//if len(matches) != len(test.results) {
		//	t.Errorf("class %s wanted %d elements, got %d instead at idx %d",
		//		test.class,
		//		len(test.results),
		//		len(matches),
		//		idx,
		//	)
		//	continue
		//}
		//for i, m := range matches {
		//	got := nodeString(m)
		if got != test.results {
			t.Errorf("class %s wanted %b, got %b instead at idx %d",
				test.class,
				test.results,
				got,
				idx,
			)
		}
		//}

		//cascadiaMatches := s.MatchAll(doc)
		//if len(matches) != len(test.results) {
		//	t.Errorf("class %s wanted %d elements, got %d instead",
		//		test.selector,
		//		len(test.results),
		//		len(matches),
		//	)
		//	continue
		//}
		//for i, m := range cascadiaMatches {
		//	got := nodeString(m)
		//	if got != test.results[i] {
		//		t.Errorf("selector %s wanted %s, got %s instead at idx %d",
		//			test.selector,
		//			test.results[i],
		//			got,
		//			idx,
		//		)
		//	}
		//}
	}
}

//func TestGetElementsByTagName(t *testing.T) {
//	var tagTests = []struct {
//		HTML     string
//		tag      string
//		selector string
//		results  []string
//	}{
//		{
//			`<p class="t1 t2">`,
//			"p",
//			"p",
//			[]string{
//				`<p class="t1 t2"></p>`,
//			},
//		},
//		{
//			`<ul class="top">
//				<li class="item" val="item1">item1</li>
//				<li class="item" val="item2">item2</li>
//				<li val="item3">
//					<ul>
//						<li class="item top" val="item3-1">item3-1</li>
//						<li val="item3-2">
//							<ul>
//								<li class="item" val="item3-2-1">item3-2-1</li>
//								<li class="top" val="item3-2-2">item3-2-2</li>
//							</ul>
//						</li>
//					</ul>
//				</li>
//				<li class="top item" val="item4">item4</li>
//				<li val="item5">item5</li>
//			</ul>`,
//			"li",
//			"li",
//			[]string{
//				`<li class="item" val="item1">item1</li>`,
//				`<li class="item" val="item2">item2</li>`,
//				`<li val="item3">
//					<ul>
//						<li class="item top" val="item3-1">item3-1</li>
//						<li val="item3-2">
//							<ul>
//								<li class="item" val="item3-2-1">item3-2-1</li>
//								<li class="top" val="item3-2-2">item3-2-2</li>
//							</ul>
//						</li>
//					</ul>
//				</li>`,
//				`<li class="item top" val="item3-1">item3-1</li>`,
//				`<li val="item3-2">
//							<ul>
//								<li class="item" val="item3-2-1">item3-2-1</li>
//								<li class="top" val="item3-2-2">item3-2-2</li>
//							</ul>
//						</li>`,
//				`<li class="item" val="item3-2-1">item3-2-1</li>`,
//				`<li class="top" val="item3-2-2">item3-2-2</li>`,
//				`<li class="top item" val="item4">item4</li>`,
//				`<li val="item5">item5</li>`,
//			},
//		},
//	}
//
//	for idx, test := range tagTests {
//		s, err := cascadia.Compile(test.selector)
//		if err != nil {
//			t.Errorf("error compiling %q: %s", test.selector, err)
//			continue
//		}
//		doc, err := html.Parse(strings.NewReader(test.HTML))
//		if err != nil {
//			t.Errorf("error parsing %q: %s", test.HTML, err)
//			continue
//		}
//
//		matches := GetElementsByTagName(doc, test.tag)
//		if len(matches) != len(test.results) {
//			t.Errorf("class %s wanted %d elements, got %d instead at idx %d",
//				test.tag,
//				len(test.results),
//				len(matches),
//				idx,
//			)
//			continue
//		}
//		for i, m := range matches {
//			got := nodeString(m)
//			if got != test.results[i] {
//				t.Errorf("class %s wanted %s, got %s instead at idx %d",
//					test.tag,
//					test.results[i],
//					got,
//					idx,
//				)
//			}
//		}
//
//		cascadiaMatches := s.MatchAll(doc)
//		if len(matches) != len(test.results) {
//			t.Errorf("class %s wanted %d elements, got %d instead",
//				test.selector,
//				len(test.results),
//				len(matches),
//			)
//			continue
//		}
//		for i, m := range cascadiaMatches {
//			got := nodeString(m)
//			if got != test.results[i] {
//				t.Errorf("selector %s wanted %s, got %s instead at idx %d",
//					test.selector,
//					test.results[i],
//					got,
//					idx,
//				)
//			}
//		}
//	}
//}
