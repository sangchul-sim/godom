package dom

import (
	"errors"
	"reflect"
	"strings"

	"fmt"

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

func GetElementsByTagName(n *html.Node, tagName string) []*html.Node {
	return getElementsByTagName(n, tagName, nil)
}

func getElementsByTagName(n *html.Node, tagName string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		if n.Data == tagName {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByTagName(c, tagName, storage)
	}
	return storage
}

func GetElementsByAttrKey(n *html.Node, key string) []*html.Node {
	return getElementsByAttrKey(n, key, nil)
}

func getElementsByAttrKey(n *html.Node, attrKey string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		if HasAttributeByKey(n, attrKey) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByAttrKey(c, attrKey, storage)
	}
	return storage
}

func GetElementsByTagNameAndClassName(n *html.Node, tagName, className string) []*html.Node {
	return getElementsByTagNameAndClassName(n, tagName, className, nil)
}

func getElementsByTagNameAndClassName(n *html.Node, tagName, className string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		if n.Data == tagName && HasClassName(n, className) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByTagNameAndClassName(c, tagName, className, storage)
	}
	return storage
}

func GetElementsByClassName(n *html.Node, className string) []*html.Node {
	return getElementsByClassName(n, className, nil)
}

func getElementsByClassName(n *html.Node, className string, storage []*html.Node) []*html.Node {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		if HasClassName(n, className) {
			storage = append(storage, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		storage = getElementsByClassName(c, className, storage)
	}
	return storage
}

func GetElementByID(n *html.Node, elementID string) (*html.Node, error) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for _, attr := range n.Attr {
			if attr.Key == AttrKeyID && attr.Val == elementID {
				return n, nil
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, err := GetElementByID(c, elementID); err == nil {
			return node, err
		}
	}
	return nil, errors.New("not found")
}

func GetAttributeByKey(n *html.Node, key string) (*html.Attribute, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return &attr, nil
		}
	}
	return nil, errors.New("not found")
}

func AddAttribute(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				if ok, _ := inArray(strings.Split(n.Attr[i].Val, " "), attr.Val); ok {
					n.Attr[i].Val += " " + attr.Val
					return
				}
			}
		}
		n.Attr = append(n.Attr, html.Attribute{
			Namespace: attr.Namespace,
			Key:       attr.Key,
			Val:       attr.Val,
		})
	}
}

func RemoveAttributeByKey(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
				break
			}
		}
	}
}

func RemoveAttributeByKeyAndVal(n *html.Node, attr *html.Attribute) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for i, _ := range n.Attr {
			if n.Attr[i].Key == attr.Key {
				n.Attr[i].Val = strings.TrimSpace(strings.Trim(n.Attr[i].Val, attr.Val))
				break
			}
		}
	}
}

func HasAttributeByKey(n *html.Node, attrKey string) (hasKey bool) {
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for _, attr := range n.Attr {
			if attr.Key == attrKey {
				hasKey = true
				return
			}
		}
	}
	return
}

func HasAttributeByKeyAndVal(n *html.Node, byAttr *html.Attribute) (hasAttr bool) {
	fmt.Println("n.Type", n.Type)
	fmt.Println("n.Attr", n.Attr)
	fmt.Println("byAttr", byAttr)
	if n.Type == html.ElementNode || n.Type == html.DocumentNode {
		for _, attr := range n.Attr {
			fmt.Println(attr.Key, attr.Val, byAttr.Key)
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

func HasClassName(n *html.Node, className string) bool {
	return HasAttributeByKeyAndVal(n, &html.Attribute{
		"", AttrKeyClass, className,
	})
}

func AddClassName(n *html.Node, className string) {
	AddAttribute(n, &html.Attribute{
		AttrKeyClass, className, "",
	})
}

func RemoveClassName(n *html.Node, className string) {
	RemoveAttributeByKeyAndVal(n, &html.Attribute{
		"", AttrKeyClass, className,
	})
}
