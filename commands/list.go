package commands

import (
	"fmt"
	"journal/core"
	"journal/crypto"
)

type ListCommand struct {
}

var cmdList = &ListCommand{}

func (l *ListCommand) Do(args []string) error {
	if len(args) != 1 {
		return ErrArgs
	}

	pass := crypto.PromptForPassword()
	defer pass.Delete()

	jrnl := new(core.Journal)
	err := jrnl.Init(args[0], &pass)
	if err != nil {
		return err
	}

	for _, entry := range jrnl.Entries {
		fmt.Printf("%s (length %d)\n", entry.Name, len(entry.Content))
	}
	return nil
}

func (i *ListCommand) Name() string {
	return "list"
}

func (i *ListCommand) Desc() string {
	return "lists all entries in the journal file"
}
func (i *ListCommand) Usage() string {
	return "journal list /path/to/journal/file"
}
func (i *ListCommand) LongUse() string {
	return ""
}
