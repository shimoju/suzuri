package suzuri

import "testing"

func TestNewClient(t *testing.T) {
	token := "accesstoken"
	client := NewClient(token)
	if client.token != token {
		t.Errorf("got %v, want %v", client.token, token)
	}
}
