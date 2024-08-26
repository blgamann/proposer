package service

import (
	"fmt"
	"net"
	"proposer/flags"
	"strconv"

	"github.com/urfave/cli/v2"
)

type CLIConfig struct {
	L1EthRpc string
	RollupRpc string
}

func checkRpcConfig(c string) error {
	_, port, err := net.SplitHostPort(c)
	if err != nil {
		return fmt.Errorf("invalid RPC config: %w", err)
	}

	p, err := strconv.Atoi(port)
	if err != nil || p < 0 || p > 65535 {
		return fmt.Errorf("invalid RPC port: %w", err)
	}

	return nil
}

func Check(c *CLIConfig) error {
	if err := checkRpcConfig(c.L1EthRpc); err != nil {
		return err
	}
	if err := checkRpcConfig(c.RollupRpc); err != nil {
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
