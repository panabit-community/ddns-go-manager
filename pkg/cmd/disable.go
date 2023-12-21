package cmd

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

type DisableCmd struct{}

func (*DisableCmd) Name() string { return "disable" }

func (*DisableCmd) Synopsis() string { return "Disable the extension." }

func (*DisableCmd) Usage() string { return "disable" }

func (p *DisableCmd) SetFlags(_ *flag.FlagSet) {}

func (p *DisableCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if doStop() == subcommands.ExitFailure {
		return subcommands.ExitFailure
	}
	if err := env.SetGlobalEnabled(false); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
