package cmd

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"

	"github.com/google/subcommands"
)

type StopCmd struct{}

func (*StopCmd) Name() string { return "stop" }

func (*StopCmd) Synopsis() string { return "Stop the extension." }

func (*StopCmd) Usage() string { return "stop" }

func (p *StopCmd) SetFlags(_ *flag.FlagSet) {}

func (p *StopCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return doStop()
}

func doStop() subcommands.ExitStatus {
	if err := ddnsgo.Stop(); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
