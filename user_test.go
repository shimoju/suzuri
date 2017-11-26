package suzuri

import (
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	userID := 1
	user, err := client.GetUser(ctx, strconv.Itoa(userID))
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}

	if user.ID != userID {
		t.Errorf("expected %v, got %v", userID, user.ID)
	}
}
