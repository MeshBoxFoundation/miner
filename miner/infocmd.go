package main

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	infoCommand = cli.Command{
		Name:        "info",
		Usage:       "miner info commands",
		Category:    "INFO COMMANDS",
		Description: `miner info commands`,
		Subcommands: []cli.Command{
			{
				Name:        "height",
				Usage:       "info height",
				Action:      utils.MigrateFlags(infoHeight),
				Description: `Get height of the blockchain for this miner.`,
			},
			{
				Name:        "in_consensus",
				Usage:       "info in_consensus",
				Action:      utils.MigrateFlags(inConsensus),
				Description: `Show if this miner is in the consensus_group.`,
			},
			{
				Name:        "name",
				Usage:       "info name",
				Action:      utils.MigrateFlags(infoName),
				Description: `Shows the name of this miner.`,
			},
			{
				Name:        "block_age",
				Usage:       "info block_age",
				Action:      utils.MigrateFlags(blockAge),
				Description: `Get age of the latest block in the chain, in seconds.`,
			},
			{
				Name:        "p2p_status",
				Usage:       "info p2p_status",
				Action:      utils.MigrateFlags(p2pStatus),
				Description: `Shows key peer connectivity status of this miner.`,
			},
			{
				Name:        "onboarding",
				Usage:       "info onboarding",
				Action:      utils.MigrateFlags(infoOnboarding),
				Description: `Get manufacturing and staking details for this miner.`,
			},
			{
				Name:        "summary",
				Usage:       "info summary",
				Action:      utils.MigrateFlags(infoSummary),
				Description: `Get a collection of key data points for this miner.`,
			},
			{
				Name:        "region",
				Usage:       "info region",
				Action:      utils.MigrateFlags(infoRegion),
				Description: `Get the operatating region for this miner.`,
			},
		},
	}
)

func infoHeight(ctx *cli.Context) error {
	return nil
}

func inConsensus(ctx *cli.Context) error {
	return nil
}

func infoName(ctx *cli.Context) error {
	return nil
}

func blockAge(ctx *cli.Context) error {
	return nil
}

func p2pStatus(ctx *cli.Context) error {
	return nil
}

func infoOnboarding(ctx *cli.Context) error {
	return nil
}

func infoSummary(ctx *cli.Context) error {
	return nil
}

func infoRegion(ctx *cli.Context) error {
	return nil
}
