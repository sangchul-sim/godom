package godom

import (
	"bytes"

	"golang.org/x/net/html"
)

type goQuery struct {
	node *html.Node
}

func NewGoQuery(n *html.Node) *goQuery {
	return &goQuery{n}
}

func (g *goQuery) GetElementsByTagName(tagName string) []*html.Node {
	return getElementsByTagName(g.node, tagName, nil)
}

func (g *goQuery) ElementsByTagName(tagName string) (queries []*goQuery) {
	for _, elem := range getElementsByTagName(g.node, tagName, nil) {
		queries = append(queries, NewGoQuery(elem))
	}
	return
}

func (g *goQuery) GetElementsByAttrKey(key string) []*html.Node {
	return getElementsByAttrKey(g.node, key, nil)
}

func (g *goQuery) ElementsByAttrKey(key string) (queries []*goQuery) {
	for _, elem := range getElementsByAttrKey(g.node, key, nil) {
		queries = append(queries, NewGoQuery(elem))
	}
	return
}

func (g *goQuery) GetElementsByAttrKeyVal(key, val string) []*html.Node {
	return getElementsByAttrKeyVal(g.node, key, val, nil)
}

func (g *goQuery) ElementsByAttrKeyVal(key, val string) (queries []*goQuery) {
	for _, elem := range getElementsByAttrKeyVal(g.node, key, val, nil) {
		queries = append(queries, NewGoQuery(elem))
	}
	return
}

func (g *goQuery) GetElementsByTagNameAndClassName(tagName, className string) []*html.Node {
	return getElementsByTagNameAndClassName(g.node, tagName, className, nil)
}

func (g *goQuery) ElementsByTagNameAndClassName(tagName, className string) (queries []*goQuery) {
	for _, elem := range getElementsByTagNameAndClassName(g.node, tagName, className, nil) {
		queries = append(queries, NewGoQuery(elem))
	}
	return
}

func (g *goQuery) GetElementsByClassName(className string) []*html.Node {
	return getElementsByClassName(g.node, className, nil)
}

func (g *goQuery) ElementsByClassName(className string) (queries []*goQuery) {
	for _, elem := range getElementsByClassName(g.node, className, nil) {
		queries = append(queries, NewGoQuery(elem))
	}
	return
}

func (g *goQuery) GetElementByID(elementID string) (*html.Node, error) {
	return getElementByID(g.node, elementID)
}

func (g *goQuery) ElementByID(elementID string) (elem *goQuery, err error) {
	node, err := getElementByID(g.node, elementID)
	if err != nil {
		return
	}
	elem = NewGoQuery(node)
	return
}

func (g *goQuery) GetAttributeByKey(key string) string {
	return getAttributeValByKey(g.node, key)
}

func (g *goQuery) AddAttribute(attr *html.Attribute) {
	addAttribute(g.node, attr)
}

func (g *goQuery) RemoveAttributeByKey(attr *html.Attribute) {
	removeAttributeByKey(g.node, attr)
}

func (g *goQuery) RemoveAttributeByKeyAndVal(attr *html.Attribute) {
	removeAttributeByKeyAndVal(g.node, attr)
}

func (g *goQuery) HasClassName(className string) bool {
	return hasAttributeByKeyAndVal(g.node, &html.Attribute{
		"", AttrKeyClass, className,
	})
}

func (g *goQuery) AddClassName(className string) {
	addAttribute(g.node, &html.Attribute{
		AttrKeyClass, className, "",
	})
}

func (g *goQuery) RemoveClassName(className string) {
	removeAttributeByKeyAndVal(g.node, &html.Attribute{
		"", AttrKeyClass, className,
	})
}

func (g *goQuery) NodeString() string {
	buf := bytes.NewBufferString("")
	html.Render(buf, g.node)
	return buf.String()
}

func (g *goQuery) Node() *html.Node {
	return g.node
}
