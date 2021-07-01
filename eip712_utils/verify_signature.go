package eip712_utils

import (
	"fmt"

	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(signature string, nonce string, publicKey string) bool {
	decodedSig, err := hexutil.Decode(signature)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return false
	}
	decodedSig[64] -= 27

	msg := []byte(fmt.Sprintf("🏆 Hi there 👋!\n\n 🎯Sign this message to prove you have access to this wallet and I’ll log you in. This won’t cost you any Ether.\n\n ✅To stop others from using your wallet, here’s a unique message ID they can’t guess:\n %s", nonce))

	prefixedNonce := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)

	hash := crypto.Keccak256Hash([]byte(prefixedNonce))
	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		fmt.Println(err)
		return false
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()
	isClientAddressEqualToRecoveredAddress := strings.ToLower(publicKey) == strings.ToLower(recoveredAddress)
	return isClientAddressEqualToRecoveredAddress
}
