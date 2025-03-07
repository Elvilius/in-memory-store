package compute

import (
	"fmt"
)

type COMMAND string

const (
	SET COMMAND = "SET"
	GET COMMAND = "GET"
	DEL COMMAND = "DEL"
)

type PreparedCommand struct {
	Cmd  COMMAND
	Key  string
	Value string
}

func findCommand(str string) (COMMAND, error) {
	switch COMMAND(str) {
	case SET:
		return SET, nil
	case GET:
		return GET, nil
	case DEL:
		return DEL, nil
	default:
		return "", fmt.Errorf("command not found")
	}
}
