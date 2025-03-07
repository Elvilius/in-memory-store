package compute

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

const paramsLen = 2

type Compute struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Compute {
	return &Compute{
		logger: logger,
	}
}

func (c *Compute) Parse(rawCmd string) (PreparedCommand, error) {
	params := strings.Fields(rawCmd)
	preparedCommand := PreparedCommand{}

	if len(params) == 0 {
		return preparedCommand, fmt.Errorf("empty command")
	}

	command, err := findCommand(params[0])
	if err != nil {
		return preparedCommand, err
	}

	preparedCommand.Cmd = command
	preparedCommand.Key = params[1]

	if len(params) > paramsLen {
		preparedCommand.Value = params[2]
	}
	return preparedCommand, nil

}
