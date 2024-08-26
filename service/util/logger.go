package util

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
)

func NewLogger() log.Logger {
	h := log.NewGlogHandler(log.NewTerminalHandler(os.Stdout, true))
	h.Verbosity(log.LevelDebug)

	return log.NewLogger(h)
}
