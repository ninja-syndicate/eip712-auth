package eip712_utils

// This function is in development.
// It signs a JSON payload as mandated by EIP712

/*
import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core"
)

// This function is in development.
// It signs a JSON payload as mandated by EIP712

func EIP712Signer(walletAddress string) error {

	// ...

	// Replace this with the address of the user's wallet
	// walletAddress := "0x61e0499cF10d341A5E45FA9c211aF3Ba9A2b50ef"
	salt, _ := GenerateNonce()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
    nonce, _ := GenerateNonce()

	signerData := signer.TypedData{
		Types: signer.Types{
			"Challenge": []signer.Type{
				{Name: "address", Type: "address"},
				{Name: "nonce", Type: "string"},
				{Name: "timestamp", Type: "string"},
			},
			"EIP712Domain": []signer.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "version", Type: "string"},
				{Name: "salt", Type: "string"},
			},
		},
		PrimaryType: "Challenge",
		Domain: signer.TypedDataDomain{
			Name:    "ETHChallenger",
			Version: "1",
			Salt:    salt,
			ChainId: math.NewHexOrDecimal256(1),
		},
		Message: signer.TypedDataMessage{
			"timestamp": timestamp,
			"address":   walletAddress,
			"nonce":     nonce,
		},
	}
	println(signerData)
	return nil
}
*/
