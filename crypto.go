package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// Crypto - crypto library
type Crypto struct{}

// NewCrypto - instantiate new Crypto
func NewCrypto() Crypto {
	return Crypto{}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Enc - encrypt data
func (c Crypto) Enc(data *[]byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, *data, nil)
	return ciphertext, nil
}

// Dec - decrypt data
func (c Crypto) Dec(data *[]byte, passphrase string) (*[]byte, error) {
	if len(*data) == 0 {
		retme := []byte("")
		return &retme, nil
	}
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := (*data)[:nonceSize], (*data)[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return &plaintext, nil
}

// EncryptFile - encrypts a file
func (c Crypto) EncryptFile(filename string, data *[]byte, passphrase string) error {
	err := os.Truncate(filename, 0)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := c.Enc(data, passphrase)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

// DecryptFile decrypts file with password
func (c Crypto) DecryptFile(filename string, passphrase string) (*[]byte, error) {
	data, _ := ioutil.ReadFile(filename)
	b, err := c.Dec(&data, passphrase)
	if err != nil {
		return nil, err
	}
	return b, nil
}
