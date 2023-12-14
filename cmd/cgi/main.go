package main

import (
	"fmt"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/html"
)

const (
	DefaultTemplatePath  = "./template"
	HttpTemplate         = "http.tpl"
	HtmlTemplate         = "html.tpl"
	PartialIndexTemplate = "index.tpl"
)

func main() {
	if err := render(); err != nil {
		os.Exit(1)
	}
}

func render() error {
	d := struct {
		Title string
	}{
		Title: "DDNS-GO Manager",
	}
	s, err := html.Render(
		DefaultTemplatePath,
		d,
		HttpTemplate, HtmlTemplate, PartialIndexTemplate,
	)
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
