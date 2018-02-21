package godom

import (
	"errors"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

const (
	AttrKeyID    = "id"
	AttrKeyClass = "class"
)

// array는 ptr로 넘기면 안됩니다.
func inArray(array interface{}, val interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func getElementsByTagName(n *html.Node, tagName string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		if n.Data == tagName {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByTagName(c, tagName, storage)
	}
	return storage
}

func getElementsByAttrKey(n *html.Node, attrKey string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		if hasAttributeByKey(n, attrKey) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByAttrKey(c, attrKey, storage)
	}
	return storage
}

func getElementsByAttrKeyVal(n *html.Node, attrKey, attrVal string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		if hasAttributeByKeyAndVal(n, &html.Attribute{"", attrKey, attrVal}) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByAttrKeyVal(c, attrKey, attrVal, storage)
	}
	return storage
}

func getElementsByTagNameAndClassName(n *html.Node, tagName, className string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		if n.Data == tagName && hasAttributeByKeyAndVal(n, &html.Attribute{
			"", AttrKeyClass, className,
		}) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByTagNameAndClassName(c, tagName, className, storage)
	}
	return storage
}

func getElementsByClassName(n *html.Node, className string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode {
		if hasAttributeByKeyAndVal(n, &html.Attribute{
			"", AttrKeyClass, className,
		}) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByClassName(c, className, storage)
	}
	return storage
}

func getElementByID(n *html.Node, elementID string) (*html.Node, error) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == AttrKeyID && attr.Val == elementID {
				return n, nil
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, err := getElementByID(c, elementID); err == nil {
			return node, err
		}
	}
	return nil, errors.New("not found")
}

func getAttributeByKey(n *html.Node, key string) (*html.Attribute, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return &attr, nil
		}
	}
	return nil, errors.New("not found")
}

func addAttribute(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				if ok, _ := inArray(strings.Split(n.Attr[i].Val, " "), attr.Val); ok {
					n.Attr[i].Val += " " + attr.Val
				} else {
					n.Attr[i].Val = attr.Val
				}
				return
			}
		}
		n.Attr = append(n.Attr, *attr)
	}
}

func removeAttributeByKey(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
				break
			}
		}
	}
}

func removeAttributeByKeyAndVal(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				n.Attr[i].Val = strings.TrimSpace(strings.Trim(n.Attr[i].Val, attr.Val))
				break
			}
		}
	}
}

func hasAttributeByKey(n *html.Node, attrKey string) (hasKey bool) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == attrKey {
				hasKey = true
				return
			}
		}
	}
	return
}

func hasAttributeByKeyAndVal(n *html.Node, byAttr *html.Attribute) (hasAttr bool) {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key != byAttr.Key {
				continue
			}
			if ok, _ := inArray(strings.Split(attr.Val, " "), byAttr.Val); ok {
				hasAttr = true
				return
			}
		}
	}
	return
}
