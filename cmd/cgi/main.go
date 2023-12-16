package main

import (
	"fmt"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
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
	if err := renderIndex(); err != nil {
		os.Exit(1)
	}
}

func renderIndex() error {
	pid, err := ddnsgo.Status()
	var status string
	if err != nil {
		status = "unknown"
	} else if pid == 0 {
		status = "inactive"
	} else {
		status = "active"
	}
	d := struct {
		ContentType string
		Title       string
		Status      string
	}{
		ContentType: "text/html; charset=GB2312",
		Title:       "DDNS-GO Manager",
		Status:      status,
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
