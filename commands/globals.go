package commands

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

var Commands = map[string]command{
	"init": cmdInit,
	"add":  cmdAdd,
	"help": cmdHelp,
	//	"serve": cmdServe,
	//	"pack":  cmdPack,
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	ErrNotJournal = errors.New("journal: can't interact with an uninitialised directory")
	ErrArgs       = errors.New("journal: invalid arguments")
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

// journal init => takes a path and creates a default sized, password protected volume in the directory given.
// journal add adds to that volume, and encrypts
// journal serve decrypts the volume and serves it over http locally only (on the specified port)
// journal pack is obselete because the whole thing is in an encrypted blob

// Difficulties:
// 		crypt,
// 		using an editor
