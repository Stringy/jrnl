package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	var fn string
	if len(args) == 1 {
		fn = args[0]
	} else {
		fn = "entry"
	}

	if !isJournalDir() {
		return ErrNotJournal
	}

	now := time.Now()
	stamp := fmt.Sprintf("%d/%d/%d/%s.md", now.Year(), now.Month(), now.Day(), fn)

	fmt.Println(stamp)
	cmd := exec.Command("vim", stamp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	if _, err := os.Stat(stamp); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File not saved, so not added to journal")
			return nil
		} else {
			return err
		}
	}

	jrnl, err := ReadJournal()
	if err != nil {
		return err
	}

	entry := Entry{Name: fn, Path: filepath.Dir(stamp)}
	jrnl.Entries = append(jrnl.Entries, entry)
	return jrnl.WriteJournal()
}

func checkOrCreateDirs(now time.Time) error {
	yrdir := fmt.Sprintf("%d", now.Year())
	mndir := fmt.Sprintf("%d", now.Month())
	dydir := fmt.Sprintf("%d", now.Day())

	pth := yrdir

	if _, err := os.Stat(pth); err != nil {
		if os.IsNotExist(err) {
			direrr := os.MkdirAll(pth+"/"+mndir+"/"+dydir, os.FileMode(0700))
			return direrr
		} else {
			return err
		}
	}

	pth += "/" + mndir
	if _, err := os.Stat(pth); err != nil {
		if os.IsNotExist(err) {
			direrr := os.MkdirAll(pth+"/"+dydir, os.FileMode(0700))
			return direrr
		} else {
			return err
		}
	}

	pth += "/" + dydir
	if _, err := os.Stat(pth); err != nil {
		if os.IsNotExist(err) {
			direrr := os.MkdirAll(pth, os.FileMode(0700))
			return direrr
		} else {
			return err
		}
	}

	return nil
}

func (i *AddCommand) Name() string    { return i.name }
func (i *AddCommand) Desc() string    { return i.desc }
func (i *AddCommand) Usage() string   { return i.usage }
func (i *AddCommand) LongUse() string { return i.longuse }
