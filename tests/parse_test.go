package tests

import (
	"bytes"

	"golang.org/x/net/html"
)

func nodeString(n *html.Node) string {
	buf := bytes.NewBufferString("")
	html.Render(buf, n)
	return buf.String()
}
