package main

import (
	"context"
	"flag"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/cmd"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

func init() {
	env.Init()
}

func main() {
	subcommands.Register(&cmd.DisableCmd{}, "")
	subcommands.Register(&cmd.EnableCmd{}, "")
	subcommands.Register(&cmd.StartCmd{}, "")
	subcommands.Register(&cmd.StatusCmd{}, "")
	subcommands.Register(&cmd.StopCmd{}, "")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
