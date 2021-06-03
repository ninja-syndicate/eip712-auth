package eip712_utils

import "testing"

func TestVerifySignature(t *testing.T) {
	verificationResult := VerifySignature(
		"0x597bdaa99f03999c1e89eba6fe94c72fa8fb14ab423cf2e9b5fb52b1702d5f094d0025429b9855f25c193bdc4cb4007c69bc7abd4fbc668d452a62ca5d478e511b",
		"9f68d6688254713cc648fccbf42ed510e19b40337f6fcf53a286b5e23b7c3f97",
		"0x016e60506b80ad835f13048157967cea06da057e",
	)
	expectedResult := true
	if expectedResult != verificationResult {
		t.Errorf("Expected is NOT EQUAL to Result")
	}
}
