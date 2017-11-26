package suzuri

import (
	"context"
	"net/http"
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()

	userID := 7
	endpoint := "/users/" + strconv.Itoa(userID)

	stub.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		expected := "GET"
		actual := r.Method
		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		http.ServeFile(w, r, "testdata/users-7.json")
	})

	user, err := client.GetUser(ctx, userID)
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}

	expected := userID
	actual := user.ID
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	cancelCtx, cancel := context.WithCancel(ctx)
	cancel()
	user, err = client.GetUser(cancelCtx, userID)
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
}
