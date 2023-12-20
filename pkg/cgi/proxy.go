package cgi

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type ProxyCmd struct{}

func (*ProxyCmd) Name() string { return "proxy" }

func (*ProxyCmd) Synopsis() string { return "Proxy ddns-go apis." }

func (*ProxyCmd) Usage() string { return "proxy <API>" }

func (p *ProxyCmd) SetFlags(_ *flag.FlagSet) {}

func (p *ProxyCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	s, err := Request(f.Arg(0))
	if err != nil {
		fmt.Printf("proxy failed: %v", err)
		return subcommands.ExitFailure
	}
	fmt.Println(s)
	return subcommands.ExitSuccess
}
