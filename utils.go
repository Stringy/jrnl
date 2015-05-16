package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/howeyc/gopass"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

type Journal struct {
	Entries []JournalEntry
	nonce   []byte
}

func (j *Journal) SetNonce(n []byte) {
	j.nonce = make([]byte, 12)
	copy(j.nonce, n)
}

type JournalEntry struct {
	Name    string
	Content string
}

type Password struct {
	pwd []byte
}

func (p *Password) Clear() {
	for i := range pwd {
		pwd[i] = 0
	}
}

func (p *Password) Passwd() []byte {
	return p.pwd
}

func ReadJournal(path string, pass *Password) (*Journal, error) {

	buf, err := getJournalFile(path)
	if err != nil {
		return nil, err
	}

	// buf[:12] == nonce
	// buf[12:] == data

	dec, err := decrypt(buf[12:], buf[:12], getKey())
	if err != nil {
		return nil, err
	}

	jrnl := new(Journal)
	err = json.Unmarshal(dec, jrnl)
	if err != nil {
		return nil, err
	}

	jrnl.SetNonce(buf[:12])
	return jrnl, nil
}

func WriteJournal(jrnl *Journal, path string) error {
	buf, err := json.Marshal(jrnl)
	if err != nil {
		return err
	}

	enc, err := encrypt(buf, jrnl.GetNonce(), getKey())
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(jrnl.GetNonce())
	if err != nil {
		return err
	}

	_, err = f.Write(enc)
	if err != nil {
		return err
	}
	return nil
}

func getJournalFile(path string) ([]byte, error) {
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

func getKey() []byte {
	fmt.Printf("Password: ")
	pass := gopass.GetPasswdMasked()
	return pass
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
