package cgi

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
)

const (
	cgiPrefix = "CGI_"
)

var (
	params = map[string][]string{
		"save": {
			"NotAllowWanAccess",
			"Username",
			"Password",
			"WebhookURL",
			"WebhookRequestBody",
			"WebhookHeaders",
			"DnsConf",
		},
	}
)

func Request(api string) (string, error) {
	if len(api) == 0 {
		r, err := get(api)
		if err != nil {
			return "", err
		}
		if err := intercept(r, os.Stdout); err != nil {
			return "", err
		}
	}
	switch ParseMethod() {
	case "get":
		return get(api)
	case "post":
		return post(api)
	}
	return "", nil
}

func get(api string) (string, error) {
	p := url.Values{}
	for _, v := range params[api] {
		p.Add(v, ParseParameter(v))
	}
	resp, err := http.Get(ddnsgo.Url + "/" + api)
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

func post(api string) (string, error) {
	p := url.Values{}
	for _, v := range params[api] {
		p.Add(v, ParseParameter(v))
	}
	resp, err := http.PostForm(ddnsgo.Url+"/"+api, p)
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

func ParseAction() string {
	return ParseParameter("action")
}

func ParseMethod() string {
	m := os.Getenv("REQUEST_METHOD")
	if len(m) == 0 {
		return "get"
	}
	return strings.ToLower(m)
}

func ParseParameter(key string) string {
	return os.Getenv(cgiPrefix + key)
}
