package service

import (
	"fmt"
	"proposer/flags"

	"github.com/urfave/cli/v2"
)

func New() cli.ActionFunc {
	return func(c *cli.Context) error {
		if err := flags.CheckRequired(c); err != nil {
			return err
		}

		logger := NewLogger()

		cfg := NewConfig(c)
		if err := Check(cfg); err != nil {
			return fmt.Errorf("invalid CLI flags: %w", err)
		}

		logger.Info("Initializing L2Output Submitter", "version", c.App.Version)

		return nil
	}
}
