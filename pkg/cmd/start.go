package cmd

import (
	"context"
	"flag"
	"os"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"
	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

type StartCmd struct{}

func (*StartCmd) Name() string { return "start" }

func (*StartCmd) Synopsis() string { return "Start the extension." }

func (*StartCmd) Usage() string { return "start" }

func (p *StartCmd) SetFlags(_ *flag.FlagSet) {}

func (p *StartCmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return doStart()
}

func doStart() subcommands.ExitStatus {
	if err := os.Chmod(env.DdnsGoBinary, 0755); err != nil {
		return subcommands.ExitFailure
	}
	if err := env.CopyDir(env.ExtensionCgiStorageDir, env.ExtensionCgiDir, 0755); err != nil {
		return subcommands.ExitFailure
	}
	if err := env.CopyDir(env.ExtensionWebTemplatesStorageDir, env.ExtensionWebTemplatesDir, 0644); err != nil {
		return subcommands.ExitFailure
	}
	if len(ddnsgo.Username()) == 0 || len(ddnsgo.Password()) == 0 {
		if err := ddnsgo.GenerateCredentials(); err != nil {
			return subcommands.ExitFailure
		}
	}
	if err := ddnsgo.Stop(); err != nil {
		return subcommands.ExitFailure
	}
	if _, err := ddnsgo.Start(ddnsgo.DefaultStartOpts); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
