package commands

import (
	"fmt"
	"io/ioutil"
	"journal/core"
	"journal/crypto"
	"os"
	"os/exec"
	"time"
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

	b, err := vipe()
	if err != nil {
		return err
	}

	now := time.Now()
	name := fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())

	jrnl.Entries = append(jrnl.Entries, core.JournalEntry{Name: name, Content: string(b)})
	return jrnl.Save()
}

// vipe will act as the intermediary between $EDITOR and this program.
// based on https://github.com/madx/moreutils/blob/master/vipe
func vipe() ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	file, err := ioutil.TempFile("/tmp/", "vipe_")
	if err != nil {
		return nil, err
	}
	defer func() {
		file.Close()
		if err := os.Remove(file.Name()); err != nil {
			panic(err)
		}
	}()

	//fmt.Println(file.Name())

	cmd := exec.Command(editor, file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (i *AddCommand) Name() string    { return i.name }
func (i *AddCommand) Desc() string    { return i.desc }
func (i *AddCommand) Usage() string   { return i.usage }
func (i *AddCommand) LongUse() string { return i.longuse }
