package main

import ()

type PackCommand struct {
	name    string
	desc    string
	usage   string
	longuse string
}

var cmdPack = &PackCommand{
	name:    "pack",
	desc:    "packs and encrypts the journal",
	usage:   "",
	longuse: ``,
}

func (p *PackCommand) Do(args []string) error {
	return nil
}

func (p *PackCommand) Name() string    { return p.name }
func (p *PackCommand) Desc() string    { return p.desc }
func (p *PackCommand) Usage() string   { return p.usage }
func (p *PackCommand) LongUse() string { return p.longuse }
