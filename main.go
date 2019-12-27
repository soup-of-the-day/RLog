package main

import (
	"fmt"
	"os"
	. "./commands"
)

func main() {
	command, err := CreateCommand(os.Args)
	if err != nil {
		panic(err)
	}
	err = command.CheckArgs()
	if err != nil {
		fmt.Println("----- Incorrect program arguments -----")
		fmt.Println(command.Usage())
		os.Exit(1)
	}
	err = command.Execute()
	if err != nil {
		panic(err)
	}
}