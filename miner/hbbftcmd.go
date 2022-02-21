package main

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	hbbftCommand = cli.Command{
		Name:        "hbbft",
		Usage:       "miner hbbft commands",
		Category:    "HBBFT COMMANDS",
		Description: `miner hbbft commands`,
		Subcommands: []cli.Command{
			{
				Name:        "status",
				Usage:       "hbbft status",
				Action:      utils.MigrateFlags(hbbftStatus),
				Description: `Display hbbft status.`,
			},
			{
				Name:        "queue",
				Usage:       "hbbft queue",
				Action:      utils.MigrateFlags(hbbftQueue),
				Description: `Display hbbft message queue.`,
			},
			{
				Name:        "skip",
				Usage:       "hbbft skip",
				Action:      utils.MigrateFlags(hbbftSkip),
				Description: `Skip current hbbft round.`,
			},
			{
				Name:        "group",
				Usage:       "hbbft group",
				Action:      utils.MigrateFlags(hbbftGroup),
				Description: `Display current hbbft group.`,
			},
			{
				Name:        "perf",
				Usage:       "hbbft perf",
				Action:      utils.MigrateFlags(hbbftPerf),
				Description: `Show performance of current group members.`,
			},
		},
	}
)

func hbbftStatus(ctx *cli.Context) error {
	return nil
}

func hbbftQueue(ctx *cli.Context) error {
	return nil
}

func hbbftSkip(ctx *cli.Context) error {
	return nil
}

func hbbftGroup(ctx *cli.Context) error {
	return nil
}

func hbbftPerf(ctx *cli.Context) error {
	return nil
}
