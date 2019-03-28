package main

import (
	"context"
	"github.com/docker/docker/client"
)

var cli *client.Client
var ctx context.Context
var err error

func main() {
	ctx = context.Background()
	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)
	listContainer()

	for {
		print()
	}
}