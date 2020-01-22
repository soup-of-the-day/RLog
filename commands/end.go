package commands

import (
	"../practice_log"
	"strconv"
)

// Command for ending a practice session and updating the csv file
type EndCommand struct {
	commandArgs []string
}

// Update a practice log with the information available in this command's arguments
func (e *EndCommand) updatePLogEntries(p *practice_log.PracticeLog) error {
	var bpm int64
	var err error
	for len(e.commandArgs) > 0 {
		bpm, err = strconv.ParseInt(e.commandArgs[1], 10, 64)
		if err != nil {
				return err
		}
		p.UpdateEntry(e.commandArgs[0], int(bpm))
		e.commandArgs = e.commandArgs[2:]
	}

	return nil
}

// Update the csv with commandArgs and hope it is formatted correctly I guess
func (e *EndCommand) Execute() error {
	plog, err := practice_log.ParseRudiments("rudiments.csv")
	if err != nil {
		return err
	}
	err = e.updatePLogEntries(plog)
	if err != nil {
		return err
	}
	err = practice_log.SaveRudiments(plog, "rudiments.csv")
	if err != nil {
		return err
	}

	return nil
}

func (c *EndCommand) Usage() string {
	return "end [ (<3-letter abbreviation of rudiment>, <BPM you attained>)  .... ]"
}

// Reject arguments that aren't handed in as a list of tuples (<abbrev>,<bpm>)
func (c *EndCommand) CheckArgs() error {
	return nil
}

