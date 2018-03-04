package godom

import (
	"bytes"

	"golang.org/x/net/html"
)

type goDom struct {
	node *html.Node
}

func NewGoDom(n *html.Node) *goDom {
	return &goDom{n}
}

func (g *goDom) GetElementsByTagName(tagName string) []*html.Node {
	return getElementsByTagName(g.node, tagName, nil)
}

func (g *goDom) GetElementsByAttrKey(key string) []*html.Node {
	return getElementsByAttrKey(g.node, key, nil)
}

func (g *goDom) GetElementsByAttrKeyVal(key, val string) []*html.Node {
	return getElementsByAttrKeyVal(g.node, key, val, nil)
}

func (g *goDom) GetElementsByTagNameAndClassName(tagName, className string) []*html.Node {
	return getElementsByTagNameAndClassName(g.node, tagName, className, nil)
}

func (g *goDom) GetElementsByClassName(className string) []*html.Node {
	return getElementsByClassName(g.node, className, nil)
}

func (g *goDom) GetElementByID(elementID string) (*html.Node, error) {
	return getElementByID(g.node, elementID)
}

func (g *goDom) GetInnerText() string {
	return getInnerText(g.node)
}

func (g *goDom) GetAttributeByKey(key string) *html.Attribute {
	return getAttributeValByKey(g.node, key)
}

func (g *goDom) AddAttribute(attr *html.Attribute) {
	addAttribute(g.node, attr)
}

func (g *goDom) RemoveAttributeByKey(attr *html.Attribute) {
	removeAttributeByKey(g.node, attr)
}

func (g *goDom) RemoveAttributeByKeyAndVal(attr *html.Attribute) {
	removeAttributeByKeyAndVal(g.node, attr)
}

func (g *goDom) HasClassName(className string) bool {
	return hasAttributeByKeyAndVal(g.node, &html.Attribute{
		"", AttrKeyClass, className,
	})
}

func (g *goDom) AddClassName(className string) {
	addAttribute(g.node, &html.Attribute{
		AttrKeyClass, className, "",
	})
}

func (g *goDom) RemoveClassName(className string) {
	removeAttributeByKeyAndVal(g.node, &html.Attribute{
		"", AttrKeyClass, className,
	})
}

func (g *goDom) NodeString() string {
	buf := bytes.NewBufferString("")
	html.Render(buf, g.node)
	return buf.String()
}

func (g *goDom) Node() *html.Node {
	return g.node
}
