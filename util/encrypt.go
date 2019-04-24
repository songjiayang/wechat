package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
)

func DecryptedData(sessionKey, iv, encryptedData string, bindData interface{}) error {
	_sessionKey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return err
	}

	_iv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return err
	}

	_encryptedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return err
	}

	plainText, err := AesCBCDncrypt(_encryptedData, _sessionKey, _iv)
	if err != nil {
		return err
	}

	return json.Unmarshal(plainText, &bindData)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	return origData[:(length - int(origData[length-1]))]
}

func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		errors.New("ciphertext too short")
	}

	if len(encryptData)%blockSize != 0 {
		errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptData, encryptData)

	return PKCS7UnPadding(encryptData), nil
}
