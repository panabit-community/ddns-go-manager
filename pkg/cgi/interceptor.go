package interceptor

import (
	gohtml "golang.org/x/net/html"
	"io"
	"strings"
)

func intercept(html string, output io.Writer) error {
	h, err := gohtml.Parse(strings.NewReader(html))
	if err != nil {
		return err
	}
	edit(h)
	if err := gohtml.Render(output, h); err != nil {
		return err
	}
	return nil
}

func edit(root *gohtml.Node) {
	var queue []*gohtml.Node
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Type == gohtml.ElementNode {
			for _, attr := range node.Attr {
				pref := "./static/"
				if attr.Key == "href" && strings.HasPrefix(attr.Val, pref) {
					attr.Val = "/html/App/ddns-go/static/" + attr.Val[len(pref):]
				}
				if attr.Key == "src" && strings.HasPrefix(attr.Val, pref) {
					attr.Val = "/html/App/ddns-go/static/" + attr.Val[len(pref):]
				}
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			queue = append(queue, c)
		}
	}
}
