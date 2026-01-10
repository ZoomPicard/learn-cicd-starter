package auth

import (
	"net/http"
	"testing"
)

func TestValidGetAPIKey(t *testing.T) {

	type test struct {
		input string
	}

	tests := []test{
		{input: "TestKey"},
		{input: "Bob"},
		{input: ""},
	}

	for _, tc := range tests {
		testHeader := http.Header{"Authorization": {"ApiKey " + tc.input}}
		value, _ := GetAPIKey(testHeader)
		if value != tc.input {
			t.Fatalf("expected: %v, got: %v", tc.input, value)
		}
	}
}

func TestNoAuthHeader(t *testing.T) {
	testHeader := make(http.Header)
	_, err := GetAPIKey(testHeader)
	if err.Error() != "no authorization header included" {
		t.Fatal("Did not get the expected no authorization header error")
	}
}

func TestMalformedAuthHeader(t *testing.T) {
	testHeader := http.Header{"Authorization": {"DifferentAuth"}}
	_, err := GetAPIKey(testHeader)
	if err.Error() != "malformed authorization heade" {
		t.Fatal("Did not get the expected malformed authorization header error")
	}
}
