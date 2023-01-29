package api_access

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gavin-xin/paraswap/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

var backendAddress = ""
var accessTime int64 = 1000
var accessStep = time.Second

func Action(ctx *cli.Context) error {
	client, err := ethclient.Dial(backendAddress)
	if err != nil {
		return err
	}

	fmt.Println("success create connection")

	wait.NonSlidingUntilWithContext(context.Background(), func(ctx context.Context) {
		for i := int64(0); i < accessTime; i++ {
			apiAccessQPS.WithLabelValues("meganode", "eth_getBlockByNumber").Inc()
			block, err := client.BlockByNumber(context.Background(), nil)
			if err != nil {
				logger.GetLogger().Error("request failed", zap.Error(err))
				apiAccessErr.WithLabelValues("meganode", "eth_getBlockByNumber", err.Error()).Inc()
			} else {
				logger.GetLogger().Debug("get latest block", zap.String("block_hash", block.Header().Hash().String()))
			}

		}
	}, accessStep)

	return nil
}

func NewCommand() *cli.Command {

	return &cli.Command{
		Name:    "Mega Node API Access",
		Aliases: []string{"api"},
		Usage:   "Access Mega Node API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "backend-address",
				Value:       "http://localhost:8545",
				Usage:       "backend address",
				Destination: &backendAddress,
				Required:    true,
			},
			&cli.Int64Flag{
				Name:        "access-times",
				Value:       1000,
				Usage:       "access times",
				Destination: &accessTime,
			},
			&cli.DurationFlag{
				Name:        "access-step",
				Value:       time.Second,
				Usage:       "access step",
				Destination: &accessStep,
			},
		},
		Action: Action,
	}
}
