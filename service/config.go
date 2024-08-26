package service

import (
	"proposer/flags"

	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	L1EthRpc string
	RollupRpc string
}

func checkRpcConfig(c string) error {
	// todo: check rpc config

	return nil
}

func Check(cliCtx *CLIConfig) error {
	if err := checkRpcConfig(cliCtx.L1EthRpc); err != nil {
		return err
	}
	if err := checkRpcConfig(cliCtx.RollupRpc); err != nil {
		return err
	}

	return nil
}

func NewConfig(c *cli.Context) *CLIConfig {
	return &CLIConfig{
		L1EthRpc: c.String(flags.L1EthRpcFlag.Name),
		RollupRpc: c.String(flags.RollupRpcFlag.Name),
	}
}
