package cmd

import (
	"context"
	"flag"
	"fmt"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

type StatusCmd struct{}

func (*StatusCmd) Name() string { return "status" }

func (*StatusCmd) Synopsis() string { return "Get status of the extension." }

func (*StatusCmd) Usage() string { return "status" }

func (p *StatusCmd) SetFlags(_ *flag.FlagSet) {}

func (p *StatusCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var s string
	if env.GlobalEnabled() {
		s = "enable"
	} else {
		s = "disable"
	}
	fmt.Println(s)
	return subcommands.ExitSuccess
}
