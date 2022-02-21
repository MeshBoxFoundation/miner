package main

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	dkgCommand = cli.Command{
		Name:        "dkg",
		Usage:       "miner dkg commands",
		Category:    "DKG COMMANDS",
		Description: `miner dkg commands`,
		Subcommands: []cli.Command{
			{
				Name:        "status",
				Usage:       "dkg status",
				Action:      utils.MigrateFlags(dkgStatus),
				Description: `Display dkg status.`,
			},
			{
				Name:        "queue",
				Usage:       "dkg queue",
				Action:      utils.MigrateFlags(dkgQueue),
				Description: `Display dkg queue.`,
			},
			{
				Name:        "running",
				Usage:       "dkg running",
				Action:      utils.MigrateFlags(dkgRunning),
				Description: `Display the running dkg group on a non-dkg mode.`,
			},
			{
				Name:        "next",
				Usage:       "dkg next",
				Action:      utils.MigrateFlags(dkgNext),
				Description: `Display the next block at which an election will take place.`,
			},
		},
	}
)

func dkgStatus(ctx *cli.Context) error {
	return nil
}

func dkgQueue(ctx *cli.Context) error {
	return nil
}

func dkgRunning(ctx *cli.Context) error {
	return nil
}

func dkgNext(ctx *cli.Context) error {
	return nil
}
