package main

import (
	"context"
	"flag"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cmd/start"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cmd/status"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&start.Cmd{}, "")
	subcommands.Register(&status.Cmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
