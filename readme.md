## EIP712 Auth

EIP-712: Ethereum typed structured data hashing and signing. Link: [EIP712](https://eips.ethereum.org/EIPS/eip-712)

## What are the project's objectives?

- To implement a go server that authenticates a user via [Metamask](https://metamask.io/)

## Programming Languages

[Golang](https://golang.org/)

## Domain

- Blockchain development
- Cryptography

## How to test?
- Clone the repo & cd into it
- Hit `go run main.go`. Your server starts running at `localhost:8080`
- Navigate to `http://localhost:8080/static/index.html`
- If `metamask` in not loaded, it will prompt the metamask extension as soon as the page loads
- Click on `getNonce`. This populates the nonce sent by the server
- Click on `sign`. This will prompt metamask popup to sign.
- Click on `verifySignature`. This will prompt an alert depending on the API response.
