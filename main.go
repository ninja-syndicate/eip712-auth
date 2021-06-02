package main

import (
	"bytes"
<<<<<<< HEAD
=======

>>>>>>> feat: create formatted string payload to sign
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/ninja-software/terror/v2"

	"github.com/go-chi/chi/v5"

	EIP712Sign "eip712-auth/eip712_utils"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	r.Post("/request_nonce", WithError(RequestNonceHandler()))
	r.Post("/signed_message", WithError(SignedMessageHandler()))
	fmt.Println("Starting server on :8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

// WithError wraps a route handler with an error handling
func WithError(next func(w http.ResponseWriter, r *http.Request) (int, error)) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		code, err := next(w, r)
		if err != nil {
			terror.Echo(err)
			http.Error(w, fmt.Sprintf(`{"message":"%s"}`, terror.Error(err, "").Error()), code)
			return
		}
	}
	return fn
}

func RequestNonceHandler() func(w http.ResponseWriter, r *http.Request) (int, error) {
	fn := func(w http.ResponseWriter, r *http.Request) (int, error) {
		nonce, err := EIP712Sign.GenerateNonce()
		if err != nil {
			return http.StatusInternalServerError, err
		}
		w.Write([]byte(nonce))
		return http.StatusAccepted, err
	}
	return fn
}

func SignedMessageHandler() func(w http.ResponseWriter, r *http.Request) (int, error) {
	fn := func(w http.ResponseWriter, r *http.Request) (int, error) {
		return http.StatusNotImplemented, terror.ErrNotImplemented
	}
	return fn
}

func Verify(userAddress common.Address, incomingSignature string, fileID string) (bool, error) {
	fmt.Println("userAddress", userAddress)
	fmt.Println("incomingSignature", incomingSignature)
	fmt.Println("fileID", fileID)
	signature, err := hex.DecodeString(incomingSignature[2:])
	if err != nil {
		return false, terror.Error(err)
	}
	if len(signature) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(signature))
	}

	if signature[64] != 27 && signature[64] != 28 {
		return false, fmt.Errorf("invalid recovery id: %d", signature[64])
	}
	signature[64] -= 27

	h := crypto.Keccak256Hash([]byte(fileID))
	pubKeyRaw, err := crypto.Ecrecover(h.Bytes(), signature)
	if err != nil {
		return false, fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return false, terror.Error(err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(userAddress.Bytes(), recoveredAddr.Bytes()) {
		return false, terror.Error(fmt.Errorf("addresses do not match: %s vs %s", userAddress, recoveredAddr))
	}
	return true, nil
}
