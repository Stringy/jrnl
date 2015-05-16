package main

import (
	"os"
)

func main() {

	getDirStructure()

	if len(os.Args) < 2 {
		cmdHelp.Do([]string{})
		os.Exit(0)
	}

	if _, ok := cmds[os.Args[1]]; ok {
		err := cmds[os.Args[1]].Do(os.Args[2:])
		handle(err)
	} else {
		if os.Args[1] == "help" {
			err := cmdHelp.Do(os.Args[2:])
			handle(err)
		}
	}
}
