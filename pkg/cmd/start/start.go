package start

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

	"github.com/google/subcommands"
)

type Cmd struct{}

func (*Cmd) Name() string { return "start" }

func (*Cmd) Synopsis() string { return "Start the extension." }

func (*Cmd) Usage() string { return "start" }

func (p *Cmd) SetFlags(_ *flag.FlagSet) {}

func (p *Cmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := env.CopyDir(env.ExtensionCgiStorageDir, env.ExtensionCgiDir, 0755); err != nil {
		return subcommands.ExitFailure
	}
	if err := env.CopyDir(env.ExtensionWebTemplatesStorageDir, env.ExtensionWebTemplatesDir, 0644); err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
