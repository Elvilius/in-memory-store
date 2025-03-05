package compute

import (
	"fmt"
)

type Command string

const (
	SET Command = "SET"
	GET Command = "GET"
	DEL Command = "DEL"
)

type PreparedCommand struct {
	Cmd  Command
	Key  string
	Value string
}

func findCommand(str string) (Command, error) {
	switch Command(str) {
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
