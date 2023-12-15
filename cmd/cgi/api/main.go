package main

import (
	"fmt"
	"os"
	"strings"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
)

const (
	cgiPrefix = "CGI_"
)

func init() {
	env.Init()
}

func main() {
	var res string
	switch strings.ToLower(parseAction()) {
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

func parseAction() string {
	return parseParameter("action")
}

func parseParameter(key string) string {
	return os.Getenv(cgiPrefix + key)
}
