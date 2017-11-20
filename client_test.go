package suzuri

import (
	"context"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	token := "accesstoken"
	client := NewClient(token)
	if client.token != token {
		t.Errorf("expected %v, got %v", token, client.token)
	}

	baseURL := "https://suzuri.jp"
	if client.baseURL.String() != baseURL {
		t.Errorf("expected %v, got %v", baseURL, client.baseURL)
	}
}

func TestNewRequest(t *testing.T) {
	client := NewClient("accesstoken")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	baseURL := client.baseURL.String()
	endpoint := "/api/v1/user"
	fullURL := baseURL + endpoint

	req, err := client.newRequest(ctx, "GET", endpoint, nil)
	if err != nil {
		t.Fatalf("failed to make a new request: %v", err)
	}

	if req.URL.String() != fullURL {
		t.Errorf("expected %v, got %v", fullURL, req.URL.String())
	}
	if client.baseURL.String() != baseURL {
		t.Errorf("baseURL should not change from %v, got %v", baseURL, client.baseURL.String())
	}

	expected := "Bearer accesstoken"
	actual := req.Header.Get("Authorization")
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	expected = "application/json"
	actual = req.Header.Get("Content-Type")
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	expected = "SuzuriGo/" + version
	actual = req.Header.Get("User-Agent")
	if !strings.HasPrefix(actual, expected) {
		t.Errorf("User-Agent should start with %v, got %v", expected, actual)
	}
}
