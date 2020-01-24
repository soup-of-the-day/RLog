package commands

import (
	"errors"
	"fmt"
)

// Represents a command that our program supports.
type Command interface {
	// Should do whatever the command claims to do when called. Error if something goes wrong.
	Execute() error
	// Returns a string explaining what the command can do.
	Usage() string
	// Determines if the arguments the command will get are malformed in some way
	CheckArgs() error
}

// Creates an appropriate Command for the list of arguments. If the command doesn't exist, error
func CreateCommand(args []string) (Command, error) {
	if len(args) <= 1 {
		return nil, errors.New("Expect at least one argument: start or end\n");
	}
	command := args[1]
	switch command {
	case "start":
		return &StartCommand{args}, nil
	case "end":
		return &EndCommand{args}, nil
	}

	return nil, errors.New(fmt.Sprintf("Unsupported Command: %s", command))
}