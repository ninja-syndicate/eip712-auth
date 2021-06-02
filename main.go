package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"

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
	r.Post("/verify_signed_message", WithError(SignedMessageHandler()))
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
		nonce := string(r.Header.Get("x-api-nonce"))
		signature := string(r.Header.Get("x-api-signature"))
		publicKey := string(r.Header.Get("x-api-publickey"))

		if nonce == "" || signature == "" || publicKey == "" {
			return http.StatusBadRequest, fmt.Errorf("Nonce/signature/publicKey absent in request headers")
		}

		verificationResponse := EIP712Sign.VerifySignature(signature, nonce, publicKey)

		response := make(map[string]string)
		if verificationResponse {
			response["isSignatureValid"] = "true"
		} else {
			response["isSignatureValid"] = "false"
		}

		w.Header().Set("Content-Type", "application/json")
		if verificationResponse == false {
			response["jwtToken"] = ""
		} else {
			jwtToken, err := EIP712Sign.CreateToken(publicKey)
			if err != nil {
				log.Fatalln(err)
			}
			response["jwtToken"] = jwtToken
		}
		json.NewEncoder(w).Encode(response)
		return http.StatusAccepted, nil
	}
	return fn
}
