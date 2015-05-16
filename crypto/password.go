package crypto

import (
	"crypto/sha256"
	"fmt"
	"github.com/howeyc/gopass"
)

type Password []byte

func (p Password) Delete() {
	for i, _ := range p {
		p[i] = 0x00
	}
	for i, _ := range p {
		p[i] = 0xff
	}
	for i, _ := range p {
		p[i] = 0x00
	}
}

func PromptForPassword() Password {
	fmt.Printf("Password: ")
	pass := sha256.Sum256(gopass.GetPasswdMasked())
	return Password(pass[:])
}

func PromptForNewPassword() Password {
	var pass []byte
	for {
		fmt.Printf("Password: ")
		pass = gopass.GetPasswdMasked()
		fmt.Printf("Confirm Password: ")
		conf := gopass.GetPasswdMasked()
		if string(pass) == string(conf) {
			break
		}
		fmt.Println("Passwords don't match!")
	}
	p := sha256.Sum256(pass)
	return Password(p[:])
}
