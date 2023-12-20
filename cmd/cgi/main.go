package main

import (
	"context"
	"flag"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cgi"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

const (
	DefaultTemplatePath  = "./template"
	HttpTemplate         = "http.tpl"
	HtmlTemplate         = "html.tpl"
	PartialIndexTemplate = "index.tpl"
	RestTemplate         = "rest.tpl"
)

func init() {
	env.Init()
}

func main() {
	if act := cgi.ParseAction(); act != "" {
		callRestfulDispatcher(act)
		return
	}
	subcommands.Register(&cgi.ProxyCmd{}, "")
	flag.Parse()
	if flag.NArg() == 0 {
		if err := renderIndex(); err != nil {
			os.Exit(1)
		}
		return
	}
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
