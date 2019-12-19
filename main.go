package main

import (
	. "./practice_log"
	"fmt"
)

func main() {
	pLog, err := ParseRudiments("rudiments.csv")
	if err != nil {
		panic(fmt.Sprintf("Failed to parsed file: %s", err))
	}
	entries := pLog.GetRandomEntries(5)
	fmt.Println(" -------- Rudiments to Practice ------- ")
	for _, entry := range entries {
		fmt.Println(entry.Name, entry.BPM)
	}
}