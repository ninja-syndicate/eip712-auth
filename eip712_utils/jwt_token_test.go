package eip712_utils

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	publicKey := "0x917e60506b80ad619f13048157967cea06da057e"
	token, err := CreateToken(publicKey)
	if err != nil {
		t.Errorf("Error is %s and token is %s\n", err, token)
	}
	authorizationHeader := fmt.Sprintf("Bearer %s", token)
	x := IsTokenValid(authorizationHeader)
	if x != nil {
		t.Errorf("Error is: %v", x)
	}
}
