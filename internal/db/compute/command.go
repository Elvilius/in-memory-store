package compute

import (
	"fmt"
)

type command string

const (
	SET command = "SET"
	GET command = "GET"
	DEL command = "DEL"
)

type PreparedCommand struct {
	Cmd  command
	Key  string
	Value string
}

func findCommand(str string) (command, error) {
	switch command(str) {
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
