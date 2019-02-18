package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// Deterministically converts an arbitrary string into a 32 byte slice.
func clean(key string) []byte {
	if len(key) > 32 {
		return []byte(key[:32])
	}

	return []byte(fmt.Sprintf("%-32v", key))
}

// Encrypt encrypts the input bytes into a base64 encoded string.
func Encrypt(key string, payload []byte) (string, error) {
	block, err := aes.NewCipher(clean(key))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, payload, nil)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt attempts to decrypt the input string.
func Decrypt(key string, token string) ([]byte, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return []byte{}, err
	}

	block, err := aes.NewCipher(clean(key))
	if err != nil {
		return []byte{}, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return []byte{}, errors.New("invalid ciphertext")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}

	return plaintext, nil
}
