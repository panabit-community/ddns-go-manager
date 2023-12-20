package cgi

import (
	"context"
	"flag"

	"xie.sh.cn/panabit-ddns-go-manager/v2/pkg/ddnsgo"

	"github.com/google/subcommands"
)

type ProxyCmd struct{}

func (*ProxyCmd) Name() string { return "proxy" }

func (*ProxyCmd) Synopsis() string { return "Proxy ddns-go apis." }

func (*ProxyCmd) Usage() string { return "proxy <API>" }

func (p *ProxyCmd) SetFlags(_ *flag.FlagSet) {}

func (p *ProxyCmd) Execute(_ context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if v, ok := args[0].(string); !ok {
		return subcommands.ExitFailure
	} else {
		Request(ddnsgo.Url + v)
		return subcommands.ExitSuccess
	}
}
