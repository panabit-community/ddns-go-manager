package main

import (
	"fmt"
	"strings"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cgi"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"
)

func init() {
	env.Init()
}

func main() {
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
