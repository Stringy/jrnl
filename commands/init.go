package commands

import (
	"fmt"
	"journal/core"
	"journal/crypto"
	"os"
)

type InitCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdInit = &InitCommand{
	name:  "init",
	desc:  "initialises a journal within a directory",
	usage: "journal init /path/to/file",
	longuse: `
  journal init

  initialises a journal within the current directory 
`,
}

func (i *InitCommand) Do(args []string) error {

	if len(args) != 1 {
		fmt.Println(i.usage)
		os.Exit(1)
	}

	pass := crypto.PromptForNewPassword()
	jrnl := core.NewJournal(args[0], &pass)
	return jrnl.Save()
}

func (i *InitCommand) Name() string    { return i.name }
func (i *InitCommand) Desc() string    { return i.desc }
func (i *InitCommand) Usage() string   { return i.usage }
func (i *InitCommand) LongUse() string { return i.longuse }
