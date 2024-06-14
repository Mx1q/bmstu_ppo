package main

import (
	"context"
	"log"
	"os"
	"ppo/console/app"
	"ppo/internal/config"
	iLogger "ppo/pkg/logger"
	"ppo/pkg/storage/postgres"
)

func main() {
	ctx := context.Background()
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile(cfg.Logger.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(logFile)
	logger := iLogger.NewLogger(cfg.Logger.Level, logFile)

	db, err := postgres.NewConn(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	app.Run(db, cfg, logger)
}
