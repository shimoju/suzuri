package suzuri

import (
	"context"
	"net/url"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	token := "accesstoken"
	client := NewClient(token)
	if client.token != token {
		t.Errorf("expected %v, got %v", token, client.token)
	}

	baseURL := "https://suzuri.jp/api/v1"
	if client.baseURL.String() != baseURL {
		t.Errorf("expected %v, got %v", baseURL, client.baseURL)
	}
}

func TestSetBaseURL(t *testing.T) {
	client := NewClient("accesstoken")
	newURL := "http://localhost:8080/api"

	err := client.SetBaseURL(newURL)
	if err != nil {
		t.Fatalf("failed to set base url: %v", err)
	}

	if client.baseURL.String() != newURL {
		t.Errorf("expected %v, got %v", newURL, client.baseURL)
	}

	err = client.SetBaseURL("http://invalid url")
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}

	if client.baseURL.String() != newURL {
		t.Errorf("expected %v, got %v", newURL, client.baseURL)
	}
}

func TestNewRequest(t *testing.T) {
	client := NewClient("accesstoken")
	ctx := context.Background()
	baseURL := client.baseURL.String()
	endpoint := "/users"
	query := url.Values{}
	query.Set("name", "surisurikun")
	fullURL := baseURL + endpoint + "?" + query.Encode()

	req, err := client.newRequest(ctx, "GET", endpoint, query, nil)
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

	body := strings.NewReader(`{"text": "TEST"}`)
	req, err = client.newRequest(ctx, "POST", "/materials/text", nil, body)
	if err != nil {
		t.Fatalf("failed to make a new request: %v", err)
	}

	req, err = client.newRequest(ctx, "INVALID METHOD", endpoint, nil, nil)
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
}
