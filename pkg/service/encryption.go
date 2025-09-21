package service

import (
	"context"
	"crypto/des"
	"encoding/base64"
	"errors"
	"strings"
)

func (s *Chapa) EncryptData(ctx context.Context, secretData string) (string, error) {
	block, err := des.NewTripleDESCipher([]byte(s.Encryptionkey))
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	blockSize := 8
	padDiff := blockSize - (len(secretData) % blockSize)
	padding := strings.Repeat(string(byte(padDiff)), padDiff)
	secretData = secretData + padding

	secretDataBytes := []byte(secretData)

	if len(secretDataBytes)%blockSize != 0 {
		err := errors.New("secretData is not a multiple of the block size")
		s.logger.Error(err.Error())
		return "", err
	}

	encrypted := make([]byte, len(secretDataBytes))
	for i := 0; i < len(secretDataBytes); i += blockSize {
		block.Encrypt(encrypted[i:i+blockSize], secretDataBytes[i:i+blockSize])
	}

	encoded := base64.StdEncoding.EncodeToString(encrypted)
	return encoded, nil
}
