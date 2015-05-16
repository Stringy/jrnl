package main

import (
	"errors"
)

type command interface {
	Name() string
	Desc() string
	Usage() string
	LongUse() string
	Do([]string) error
}

var cmds = map[string]command{
	"init":  cmdInit,
	"add":   cmdAdd,
	"serve": cmdServe,
	"pack":  cmdPack,
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	ErrNotJournal = errors.New("journal: can't interact with an uninitialised directory")
)

// Ideas
//
// - Entry templating
// - Markdown rendering (html?)
// - Editing of an already created file
// - Packing of contents (into tar.gz?)
// - Encrypted eventually (probably a priority)
//
// Installation directory (environ var for root?)
// Contain man pages
// journal help [cmd]
