package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Elvilius/in-memory-store/internal/client"
	"github.com/Elvilius/in-memory-store/internal/config"
	"go.uber.org/zap"
)


func main() {
	logger, _:= zap.NewDevelopment()
	config, err := config.New()
	if err != nil {
		logger.Sugar().Fatalln(err)
	}
	scanner := bufio.NewScanner(os.Stdin)

	client, err := client.NewTCPClient(*config)
	if err != nil {
		logger.Sugar().Fatalln(err)
	}

	fmt.Print("> ")
	for scanner.Scan() {
		res, err := client.Send([]byte(scanner.Text()))
		if err != nil {
			logger.Sugar().Errorln(err)
		} else {
			fmt.Println(string(res))
		}
		fmt.Print("> ")
	}

	if scanner.Err() != nil {
		logger.Sugar().Error(scanner.Err())
		return
	}
}
