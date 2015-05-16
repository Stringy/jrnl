package main

import (
	"journal/commands"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		commands.Commands["help"].Do(os.Args)
		os.Exit(0)
	}

	if _, ok := commands.Commands[os.Args[1]]; ok {
		err := commands.Commands[os.Args[1]].Do(os.Args[2:])
		if err != nil {
			panic(err)
		}
	}
}
