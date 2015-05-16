package main

import (
	"fmt"
	"os"
)

type InitCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdInit = &InitCommand{
	name: "init",
	desc: "initialises a journal within a directory",
	longuse: ` 
  journal init

  initialises a journal within the current directory 
`,
}

func (i *InitCommand) Do(args []string) error {
	if len(args) != 0 {
		fmt.Println("Usage:", os.Args[0], "init")
		os.Exit(1)
	}

	err := os.Mkdir(".jrnl", os.FileMode(0700))
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("journal: already initialised")
		} else {
			return err
		}
	}

	return nil
}

func (i *InitCommand) Name() string    { return i.name }
func (i *InitCommand) Desc() string    { return i.desc }
func (i *InitCommand) Usage() string   { return i.usage }
func (i *InitCommand) LongUse() string { return i.longuse }
