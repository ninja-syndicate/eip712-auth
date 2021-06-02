// IGNORE THIS FILE

import (
    signer "github.com/ethereum/go-ethereum/signer/core"
    "crypto/rand"
    "time"
)
    
// ...

// Replace this with the address of the user's wallet
walletAddress := "0x61e0499cF10d341A5E45FA9c211aF3Ba9A2b50ef"
salt := "some-random-string-or-hash-here"
timestamp := strconv.FormatInt(time.Unix(), 10)

// Generate a random nonce to include in our challenge
nonceBytes := make([]byte, 32)
n, err := rand.Read(nonceBytes)
if n != 32 {
	return errors.New("nonce: n != 64 (bytes)")
} else if err != nil {
       return err
}
nonce := hex.EncodeToString(nonceBytes)

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
        ChainId: big.NewInt(1),
    },
    Message: signer.TypedDataMessage{
        "timestamp": timestamp,
        "address":   walletAddress,
        "nonce":     nonce,
    },
}