package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Journal struct {
	Entries []Entry
}

type Entry struct {
	Name string // Filename
	Path string // Path not including name
}

// reads the directory struct from the .jrnl/struct.json file
func ReadJournal() (*Journal, error) {
	if !isJournalDir() {
		return nil, ErrNotJournal
	}

	f, err := os.Open("./.jrnl/struct.json")
	// if we have no struct.json file, assume empty Journal
	if err != nil {
		if os.IsNotExists(err) {
			return &Journal{Entries: make([]Entry, 0)}, nil
		}
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	jrnl := new(Journal)
	err := json.Unmarshal(b, jrnl)
	if err != nil {
		return nil, err
	}

	return jrnl, nil
}

func (j *Journal) WriteJournal() error {
	if !isJournalDir() {
		return ErrNotJournal
	}
	f, err := os.Open("./.jrnl/struct.json")
	if err != nil {
		return err
	}
	b, err := json.Marshal(j)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func isJournalDir() bool {
	if _, err := os.Stat("./.jrnl"); err != nil {
		return os.IsNotExist(err)
	}
	return true
}
