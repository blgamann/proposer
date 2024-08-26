package util

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

const DefaultDialTimeout = 1 * time.Minute

func Dial(ctx context.Context, url string, logger log.Logger) (*ethclient.Client, error) {
	c, err := dial(url, logger)
	if err != nil {
		return nil, err
	}

	return ethclient.NewClient(c), nil
}

func dial(url string, logger log.Logger) (*rpc.Client, error) {
	client, err := rpc.Dial(url)
	if err != nil {
		logger.Error("Failed to connect to L1 client", "url", url, "err", err)
		return nil, err
	}

	return client, nil
}
