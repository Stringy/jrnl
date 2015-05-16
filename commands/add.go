package commands

import (
// "fmt"
// "os"
// "os/exec"
// "path/filepath"
// "time"
)

type AddCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdAdd = &AddCommand{
	name:    "add",
	desc:    "Add a new entry to the journal",
	longuse: ``,
}

func (a *AddCommand) Do(args []string) error {
	return nil
}

func (i *AddCommand) Name() string    { return i.name }
func (i *AddCommand) Desc() string    { return i.desc }
func (i *AddCommand) Usage() string   { return i.usage }
func (i *AddCommand) LongUse() string { return i.longuse }
