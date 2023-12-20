package main

import (
	"fmt"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/html"
)

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
		DDNSGO      struct {
			Status string
		}
	}{
		ContentType: "text/html; charset=GB2312",
		Title:       "DDNS-GO Manager",
		DDNSGO: struct {
			Status string
		}{
			Status: status,
		},
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
