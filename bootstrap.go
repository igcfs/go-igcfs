package main

import (
	"context"

	"github.com/libp2p/go-libp2p"
	host "github.com/libp2p/go-libp2p-host"
)

// Option represents command line option.
// Options are used in bootstrapping phase for the node.
type Option interface {
}

func processCmdLn() []Option {
	options := make([]Option, 0, 0)
	return options
}

func initiateHost(options []Option) (host.Host, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	// To construct a simple host with all the default settings, just use `New`
	h, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"))
	return h, cancel, err
}
