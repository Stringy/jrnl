package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func CreateNonce() ([]byte, error) {
	nonce := make([]byte, 12)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func Encrypt(plain, nonce []byte, key *Password) ([]byte, error) {

	b, err := aes.NewCipher([]byte(*key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nil, nonce, plain, nil), nil

}

func Decrypt(enc, nonce []byte, key *Password) ([]byte, error) {

	b, err := aes.NewCipher([]byte(*key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}

	return gcm.Open(nil, nonce, enc, nil)

}
