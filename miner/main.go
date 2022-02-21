package main

import (
	"fmt"
	"sort"

	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	Version = "v0.1"
	app     = cli.NewApp()
)

func init() {
	app.Action = miner
	app.Name = "miner"
	app.Usage = "the miner command line interface"
	app.Version = Version
	app.Copyright = "Copyright 2012-2030 The MeshBoxFoundation Authors"

	app.Commands = []cli.Command{
		genesisCommand,
		dkgCommand,
		hbbftCommand,
		infoCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(-1)
	}
}

func miner(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}

	return nil
}
