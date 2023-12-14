package status

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type Cmd struct{}

func (*Cmd) Name() string { return "status" }

func (*Cmd) Synopsis() string { return "Get status of the extension." }

func (*Cmd) Usage() string { return "status" }

func (p *Cmd) SetFlags(_ *flag.FlagSet) {}

func (p *Cmd) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("enable")
	return subcommands.ExitSuccess
}
