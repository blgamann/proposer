package service

import (
	"context"
	"fmt"
	"proposer/flags"
	"proposer/service/util"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

type Service struct {
	Logger log.Logger

	L1Client *ethclient.Client
	RollupClient *ethclient.Client
}

func New() cli.ActionFunc {
	return func(cliCtx *cli.Context) error {
		if err := flags.CheckRequired(cliCtx); err != nil {
			return err
		}

		cfg := NewConfig(cliCtx)
		if err := Check(cfg); err != nil {
			return fmt.Errorf("invalid CLI flags: %w", err)
		}

		var s Service

		logger := util.NewLogger()
		s.Logger = logger

		if err := s.init(cliCtx.Context, cfg, logger); err != nil {
			return fmt.Errorf("failed to initialize service: %w", err)
		}

		logger.Info("Initializing L2Output Submitter", "version", cliCtx.App.Version)
		return nil
	}
}

func (s *Service) init(ctx context.Context, cfg *CLIConfig, logger log.Logger) error {
	l1Client, err := util.Dial(ctx, cfg.L1EthRpc, logger)
	if err != nil {
		return err
	}
	s.L1Client = l1Client

	rollupClient, err := util.Dial(ctx, cfg.RollupRpc, logger)
	if err != nil {
		return err
	}
	s.RollupClient = rollupClient

	l1ChainId, _ := s.L1Client.ChainID(ctx)
	rollupChainId, _ := s.RollupClient.ChainID(ctx)
	logger.Info("Connected to L1 and Rollup", "l1ChainId", l1ChainId, "rollupChainId", rollupChainId)

	return nil
}
