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

// Creates an appropriate Command for the given list of arguments. If the command doesn't exist, return error.
func CreateCommand(args []string) (Command, error) {
	command := args[1]
	switch command {
	case "start":
		return &StartCommand{args[2:]}, nil
	}

	return nil, errors.New(fmt.Sprintf("Unsupported Command: %s", command))
}