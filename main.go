package main

// TODO - Use go modules instead of these relative imports...
import (
	"fmt"
	"os"
	. "./commands"
)

func main() {
	command, err := CreateCommand(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = command.CheckArgs()
	if err != nil {
		fmt.Println("----- Invalid program arguments -----")
		fmt.Println(err.Error())
		fmt.Println("Command usage: ", command.Usage())
		os.Exit(1)
	}
	err = command.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}