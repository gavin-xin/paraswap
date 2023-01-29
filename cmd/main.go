package main

import (
	api_access "github.com/gavin-xin/paraswap/api-access"
	"github.com/gavin-xin/paraswap/logger"
	"github.com/gavin-xin/paraswap/metrics"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

func main() {
	loggerLevel := ""
	metricsAddress := ""
	app := &cli.App{
		Commands: []*cli.Command{
			api_access.NewCommand(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "logger.level",
				Value:       "info",
				Usage:       "logger level",
				Destination: &loggerLevel,
			},
			&cli.StringFlag{
				Name:        "metrics.server",
				Value:       ":8090",
				Usage:       "prometheus metrics address",
				Destination: &metricsAddress,
			},
		},
		Before: func(context *cli.Context) error {
			level, err := zapcore.ParseLevel(loggerLevel)
			if err != nil {
				return err
				//cli.Exit(fmt.Sprintf("parse level failed, %s", err), -1)
			}
			logger.NewLogger(level)

			api_access.InitMetrics()
			go metrics.StartMetricsServer(metricsAddress)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	defer logger.GetLogger().Sync()

}
