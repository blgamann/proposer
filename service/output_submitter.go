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
	Worker *Worker

	L1Client *ethclient.Client
	RollupClient *ethclient.Client
}

func NewOutputSubmitter() cli.ActionFunc {
	return func(cliCtx *cli.Context) error {
		// required flag check
		if err := flags.CheckRequired(cliCtx); err != nil {
			return err
		}

		// validate flags
		cfg := NewConfig(cliCtx)
		if err := Check(cfg); err != nil {
			return fmt.Errorf("invalid CLI flags: %w", err)
		}

		var s Service

		logger := util.NewLogger()
		s.Logger = logger

		if err := s.init(cliCtx.Context, cfg); err != nil {
			return fmt.Errorf("failed to initialize service: %w", err)
		}
		logger.Info("Initializing L2Output Submitter", "version", cliCtx.App.Version)

		s.Worker.Start(cliCtx.Context)
		return nil
	}
}

func (s *Service) init(ctx context.Context, cfg *CLIConfig) error {
	s.initClients(ctx, cfg)
	s.initWorker(ctx, cfg, s.L1Client)

	return nil
}

func (s *Service) initClients(ctx context.Context, cfg *CLIConfig) error {
	l1Client, err := util.Dial(ctx, cfg.L1EthRpc, s.Logger)
	if err != nil {
		return err
	}
	s.L1Client = l1Client

	rollupClient, err := util.Dial(ctx, cfg.RollupRpc, s.Logger)
	if err != nil {
		return err
	}
	s.RollupClient = rollupClient

	l1ChainId, _ := s.L1Client.ChainID(ctx)
	rollupChainId, _ := s.RollupClient.ChainID(ctx)
	s.Logger.Info("Connected to L1 and Rollup", "l1ChainId", l1ChainId, "rollupChainId", rollupChainId)

	return nil
}

func (s *Service) initWorker(ctx context.Context, cfg *CLIConfig, l1Client *ethclient.Client) {
	w, err := NewWorker(ctx, cfg.L2OutputOracleAddress, l1Client, s.Logger)
	if err != nil {
		s.Logger.Error("Failed to initialize worker", "err", err)
	}

	s.Worker = w
}
