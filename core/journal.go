package core

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"journal/crypto"
	"os"
)

type Journal struct {
	Entries []JournalEntry
	pass    *crypto.Password
	nonce   []byte
	path    string
}

type JournalEntry struct {
	Name    string
	Content string
}

func NewJournal(path string, pass *crypto.Password) *Journal {
	nonce, err := crypto.CreateNonce()
	if err != nil {
		panic(err)
	}
	return &Journal{
		Entries: make([]JournalEntry, 0),
		pass:    pass,
		nonce:   nonce,
		path:    path,
	}
}

func (j *Journal) Init(path string, pass *crypto.Password) error {
	buf, err := j.loadFile(path)
	if err != nil {
		return err
	}

	dec, err := crypto.Decrypt(buf[12:], buf[:12], pass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dec, j)
	if err != nil {
		return err
	}

	copy(j.nonce, buf[:12])
	j.pass = pass
	j.path = path
	return nil
}

func (j *Journal) Save() error {

	buf, err := json.Marshal(j)
	if err != nil {
		return err
	}

	enc, err := crypto.Encrypt(buf, j.nonce, j.pass)
	if err != nil {
		return err
	}

	f, err := os.Create(j.path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	gz := gzip.NewWriter(f)
	defer gz.Close()

	_, err = gz.Write(j.nonce)
	if err != nil {
		return err
	}

	_, err = gz.Write(enc)
	if err != nil {
		return err
	}

	return nil
}

func (j *Journal) loadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	return ioutil.ReadAll(gz)
}
