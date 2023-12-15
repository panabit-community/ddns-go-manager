package main

import (
	"encoding/json"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/html"
)

const (
	DefaultTemplatePath = "./template"
	HttpTemplate        = "http.tpl"
	RestTemplate        = "rest.tpl"
)

func render(code int, data any) string {
	j, _ := json.Marshal(
		struct {
			Code int
			Data any
		}{
			Code: code,
			Data: data,
		},
	)
	d := struct {
		ContentType string
		Data        string
	}{
		ContentType: "application/json; charset=GB2312",
		Data:        string(j),
	}
	s, _ := html.Render(DefaultTemplatePath, d, HttpTemplate, RestTemplate)
	return s
}
