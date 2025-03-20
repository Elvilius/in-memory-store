package main

import (
	"context"
	"github.com/Elvilius/in-memory-store/internal/config"
	"github.com/Elvilius/in-memory-store/internal/db"
	"github.com/Elvilius/in-memory-store/internal/db/compute"
	"github.com/Elvilius/in-memory-store/internal/db/engine"
	"github.com/Elvilius/in-memory-store/internal/server"
	"go.uber.org/zap"
)

const configPath = "/Users/victor/work/in-memory-store/config.yaml"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	cfg, err := config.New()
	if err != nil {
		logger.Sugar().Fatalln(err)
	}
	engine := engine.New()
	compute := compute.New(logger)
	db := db.New(logger, engine, compute)
	server, err := server.NewTCPServer(cfg, db, logger)
	if err != nil {
		logger.Sugar().Fatalln(err)
	}

	server.Run(ctx)
}
