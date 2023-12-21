package cmd

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

type EnableCmd struct{}

func (*EnableCmd) Name() string { return "enable" }

func (*EnableCmd) Synopsis() string { return "Enable the extension." }

func (*EnableCmd) Usage() string { return "enable" }

func (p *EnableCmd) SetFlags(_ *flag.FlagSet) {}

func (p *EnableCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := env.SetGlobalEnabled(true); err != nil {
		return subcommands.ExitFailure
	}
	return doStart()
}
