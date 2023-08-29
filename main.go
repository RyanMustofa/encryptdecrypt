package encryptdecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(key string, text string) (string, error) {
	rawkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	cipherblock, err := aes.NewCipher([]byte(rawkey))
	if err != nil {
		return "", err
	}
	b64text := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b64text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(cipherblock, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b64text))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key string, text string) (string, error) {
	rawkey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}
	dectext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	btext := []byte(dectext)
	cipherblock, err := aes.NewCipher([]byte(rawkey))
	if err != nil {
		return "", err
	}
	if len(btext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := btext[:aes.BlockSize]
	btext = btext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherblock, iv)
	cfb.XORKeyStream(btext, btext)
	data, err := base64.StdEncoding.DecodeString(string(btext))
	if err != nil {
		return "", err
	}
	return string(data), nil
}
