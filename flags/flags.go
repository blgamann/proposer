package flags

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const prefix = "PROPOSER"

func makeEnvVarName(name string) []string {
	return []string{fmt.Sprintf("%s_%s", prefix, name)}
}

var (
	// Required Flags
	L1EthRpcFlag = &cli.StringFlag{
		Name:  "l1-eth-rpc",
		Usage: "L1 ETH RPC URL",
		EnvVars: makeEnvVarName("L1_ETH_RPC"),
	}
	RollupRpcFlag = &cli.StringFlag{
		Name:  "rollup-rpc",
		Usage: "HTTP provider URL for the rollup node. A comma-separated list enables the active rollup provider.",
		EnvVars: makeEnvVarName("ROLLUP_RPC"),
	}

	// Optional Flag
	L2OOAddressFlag = &cli.StringFlag{
		Name:    "l2oo-address",
		Usage:   "Address of the L2OutputOracle contract",
		EnvVars: makeEnvVarName("L2OO_ADDRESS"),
	}
)

var requiredFlags = []cli.Flag{
	L1EthRpcFlag,
	RollupRpcFlag,
}

var optionalFlags = []cli.Flag{
	L2OOAddressFlag,
}

var Flags []cli.Flag

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

func CheckRequired(c *cli.Context) error {
	for _, f := range requiredFlags {
		if !c.IsSet(f.Names()[0]){
			return fmt.Errorf("flag %s is required", f.Names()[0])
		}
	}
	return nil
}
