package cgi

import (
	"io"
	"net/http"
	"net/url"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
)

const (
	cgiPrefix = "CGI_"
)

func Request(path string, params *[]string) (string, error) {
	act := ParseAction()
	switch act {
	case "GET":
		return get(path, params)
	case "POST":
		return post(path, params)
	}
	return "", nil
}

func get(path string, params *[]string) (string, error) {
	p := url.Values{}
	for _, v := range *params {
		p.Add(v, ParseParameter(v))
	}
	resp, err := http.Get(ddnsgo.Url + path)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func post(path string, params *[]string) (string, error) {
	return "", nil
}

func ParseAction() string {
	return ParseParameter("action")
}

func ParseParameter(key string) string {
	return os.Getenv(cgiPrefix + key)
}
