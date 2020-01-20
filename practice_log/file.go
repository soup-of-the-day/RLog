// Contains logic for parsing a csv and returning a PracticeLog
package practice_log

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Parses the given csv file for rudiment entries and creates a PracticeLog of those entries
// csv format: <abbrev>,<bpm>,<name>
func ParseRudiments(filename string) (*PracticeLog, error) {
	log := make(map[string]*PracticeEntry, 0)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		bpm, err := strconv.ParseInt(line[1], 10, 64)
		log[line[0]] = &PracticeEntry{
			BPM:  int(bpm),
			Name: line[2],
		}
	}
	return NewPracticeLog(log), nil
}
