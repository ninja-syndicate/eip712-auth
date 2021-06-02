package eip712_utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func GenerateNonce() (string, error) {
	// Generate a random nonce to include in our challenge
	nonceBytes := make([]byte, 32)
	n, err := rand.Read(nonceBytes)
	if n != 32 {
		return "", errors.New("nonce: n != 64 (bytes)")
	} else if err != nil {
		return "", err
	}
	nonce := hex.EncodeToString(nonceBytes)
	return nonce, nil
}
