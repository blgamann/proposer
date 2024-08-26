package service

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
)

func NewLogger() log.Logger {
	h := log.NewGlogHandler(log.NewTerminalHandler(os.Stdout, true))

	return log.NewLogger(h)
}
