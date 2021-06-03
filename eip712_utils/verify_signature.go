package eip712_utils

import (
	"fmt"
	"log"

	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(signature string, nonce string, publicKey string) bool {
	decodedSig, err := hexutil.Decode(signature)
	if err != nil {
		log.Fatal(err)
	}
	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return false
	}
	decodedSig[64] -= 27

	companyName := "John"
	msg := []byte(fmt.Sprintf("🏆Hi! This is %s👋!\n\n 🎯Sign this message to prove you have access to this wallet and I’ll log you in. This won’t cost you any Ether.\n\n ✅To stop others from using your wallet, here’s a unique message ID they can’t guess:\n %s", companyName, nonce))

	prefixedNonce := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)

	hash := crypto.Keccak256Hash([]byte(prefixedNonce))
	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		log.Fatal(err)
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		log.Fatal(err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()
	isClientAddressEqualToRecoveredAddress := strings.ToLower(publicKey) == strings.ToLower(recoveredAddress)
	return isClientAddressEqualToRecoveredAddress
}
