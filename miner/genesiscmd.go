package main

import (
	"github.com/ethereum/go-ethereum/cmd/utils"
	"gopkg.in/urfave/cli.v1"
)

var (
	genesisCommand = cli.Command{
		Name:        "genesis",
		Usage:       "miner genesis commands",
		Category:    "GENESIS COMMANDS",
		Description: `miner genesis commands`,
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "genesis create old_genesis_file pubkey proof addrs",
				Action:      utils.MigrateFlags(createGenesis),
				Description: `Create genesis block keeping old ledger transactions.`,
			},
			{
				Name:        "forge",
				Usage:       "genesis forge pubkey key_proof addrs",
				Action:      utils.MigrateFlags(forgeGenesis),
				Description: `Create genesis block from scratch just with the addresses.`,
			},
			{
				Name:        "load",
				Usage:       "genesis load genesis_file",
				Action:      utils.MigrateFlags(loadGenesis),
				Description: `Load genesis block from file.`,
			},
			{
				Name:        "export",
				Usage:       "genesis export path",
				Action:      utils.MigrateFlags(exportGenesis),
				Description: `Write genesis block to a file.`,
			},
			{
				Name:        "key",
				Usage:       "genesis key",
				Action:      utils.MigrateFlags(keyGenesis),
				Description: `create a keypair for use as a master key.`,
			},
			{
				Name:        "proof",
				Usage:       "genesis proof privkey",
				Action:      utils.MigrateFlags(proofGenesis),
				Description: `create a key proof for adding a master key to the genesis block.`,
			},
		},
	}
)

func createGenesis(ctx *cli.Context) error {
	return nil
}

func forgeGenesis(ctx *cli.Context) error {
	return nil
}

func loadGenesis(ctx *cli.Context) error {
	return nil
}

func exportGenesis(ctx *cli.Context) error {
	return nil
}

func keyGenesis(ctx *cli.Context) error {
	return nil
}

func proofGenesis(ctx *cli.Context) error {
	return nil
}
