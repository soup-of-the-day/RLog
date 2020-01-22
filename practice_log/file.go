// Contains logic for parsing a csv and returning a PracticeLog
package practice_log

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
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

// Save the contents of a practice log as a csv file
//csv format: <abbrev>,<bpm>,<name>
func SaveRudiments(p *PracticeLog, filename string) error {
	csvString := CsvStringFromLog(p)
	err := ioutil.WriteFile(filename, []byte(csvString), 0644)
	if err != nil {
		return err
	}
	return nil;
}

// Turn a practice log into csv-formatted string
func CsvStringFromLog(p *PracticeLog) string {
	csvString := ""
	for tla, l := range p.Log {
		csvString = csvString + fmt.Sprintf("%s,%d,%s\n", tla, l.BPM, l.Name)
	}
	return csvString
}