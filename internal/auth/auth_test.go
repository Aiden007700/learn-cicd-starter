package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-api-key")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if apiKey != "test-api-key" {
		t.Errorf("expected apiKey 'test-api-key', got '%s'", apiKey)
	}
}
