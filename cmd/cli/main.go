package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Elvilius/in-memory-store/internal/db"
	"github.com/Elvilius/in-memory-store/internal/db/compute"
	"github.com/Elvilius/in-memory-store/internal/db/engine"
	"go.uber.org/zap"
)

func main() {
	logger := zap.New(nil)
	engine := engine.New()
	compute := compute.New(logger)

	db := db.New(
		logger,
		engine,
		compute,
	)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		fmt.Println(db.CommandHandle(scanner.Text()))
		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		logger.Sugar().Error(scanner.Err())
		return
	}
}
