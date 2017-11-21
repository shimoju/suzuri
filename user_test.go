package suzuri

import (
	"context"
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	client := NewClient("accesstoken")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID := 7
	user, err := client.GetUser(ctx, strconv.Itoa(userID))
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}

	if user.ID != userID {
		t.Errorf("expected %v, got %v", userID, user.ID)
	}
}
