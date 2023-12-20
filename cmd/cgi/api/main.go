package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cgi"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

func init() {
	env.Init()
}

func main() {
	if flag.NArg() == 0 {
		callRestfulDispatcher()
		return
	}
	subcommands.Register(&cgi.ProxyCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

func callRestfulDispatcher() {
	var res string
	switch strings.ToLower(cgi.ParseAction()) {
	case strings.ToLower("StartInstance"):
		res = render(startInstance())
	case strings.ToLower("StopInstance"):
		res = render(stopInstance())
	case strings.ToLower("DescribeInstanceStatus"):
		res = render(describeInstanceStatus())
	default:
	}
	fmt.Println(res)
}
