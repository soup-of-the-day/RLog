package commands

import (
	. "../practice_log"
	"errors"
	"fmt"
	"strconv"
)

// Command for starting a practice session.
type StartCommand struct {
	commandArgs []string
}

// Should start a practice session by printing out some random number of rudiments and their respective BPMs.
func (c *StartCommand) Execute() error {
	pLog, err := ParseRudiments("rudiments.csv")
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to parsed file: %s", err))
	}
	numEntries, _ := strconv.ParseInt(c.commandArgs[0], 10, 0)
	entries := pLog.GetRandomEntries(int(numEntries))
	fmt.Println(" -------- Rudiments to Practice ------- ")
	for _, entry := range entries {
		fmt.Println(entry.Name, entry.BPM)
	}

	return nil
}

func (c *StartCommand) Usage() string {
	return "start <number of Rudiments to practice>"
}

// This command should only accept a single integer indicating the number of rudiments to practice
func (c *StartCommand) CheckArgs() error {
	// Only break if first arg can't be turned into an int
	_, err := strconv.ParseInt(c.commandArgs[0], 10, 0)
	if err != nil {
		return err
	}
	return nil
}

