package main

import (
	"fmt"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/html"
)

const (
	DefaultTemplatePath  = "./template"
	HttpTemplate         = "http.tpl"
	HtmlTemplate         = "html.tpl"
	PartialIndexTemplate = "index.tpl"
)

func init() {
	env.Init()
}

func main() {
	if err := render(); err != nil {
		os.Exit(1)
	}
}

func render() error {
	d := struct {
		ContentType string
		Title       string
	}{
		ContentType: "text/html; charset=GB2312",
		Title:       "DDNS-GO Manager",
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
