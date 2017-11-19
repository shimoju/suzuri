package suzuri

import "testing"

func TestNewClient(t *testing.T) {
	token := "accesstoken"
	client := NewClient(token)
	if client.token != token {
		t.Errorf("expected %v, got %v", token, client.token)
	}

	baseURL := "https://suzuri.jp"
	if client.url.String() != baseURL {
		t.Errorf("expected %v, got %v", baseURL, client.url)
	}
}
