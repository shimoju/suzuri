package suzuri

import (
	"context"
	"net/http"
	"testing"
)

func TestListItems(t *testing.T) {
	setup()
	defer teardown()

	stub.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		expected := "GET"
		actual := r.Method
		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}

		http.ServeFile(w, r, "testdata/items.json")
	})

	items, err := client.ListItems(ctx)
	if err != nil {
		t.Fatalf("failed to get items: %v", err)
	}

	expected := 4
	actual := len(items)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	expected = 6
	actual = len(items[2].Angles)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	expected = 0
	actual = len(items[3].Angles)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	expected = 4
	actual = len(items[0].Variants)
	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	cancelCtx, cancel := context.WithCancel(ctx)
	cancel()
	items, err = client.ListItems(cancelCtx)
	if err == nil {
		t.Errorf("should return error, got %v", err)
	}
}
