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

func (g *goQuery) GetElementsByAttrKey(key string) []*html.Node {
	return getElementsByAttrKey(g.node, key, nil)
}

func (g *goQuery) GetElementsByAttrKeyVal(key, val string) []*html.Node {
	return getElementsByAttrKeyVal(g.node, key, val, nil)
}

func (g *goQuery) GetElementsByTagNameAndClassName(tagName, className string) []*html.Node {
	return getElementsByTagNameAndClassName(g.node, tagName, className, nil)
}

func (g *goQuery) GetElementsByClassName(className string) []*html.Node {
	return getElementsByClassName(g.node, className, nil)
}

func (g *goQuery) GetElementByID(elementID string) (*html.Node, error) {
	return getElementByID(g.node, elementID)
}

func (g *goQuery) GetAttributeByKey(key string) (*html.Attribute, error) {
	return getAttributeByKey(g.node, key)
}

func (g *goQuery) AddAttribute(attr *html.Attribute) {
	addAttribute(g.node, attr)
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
