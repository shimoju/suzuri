package suzuri

import (
	"context"
)

// Items is a collection of Item.
type Items struct {
	Items []Item `json:"items"`
}

// Item is a product type you can create in SUZURI.
type Item struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Angles       []string      `json:"angles"`
	HumanizeName string        `json:"humanizeName"`
	Variants     []ItemVariant `json:"variants"`
}

// ItemVariant is a combination of Color and Size.
type ItemVariant struct {
	ID        int       `json:"id"`
	Price     int       `json:"price"`
	Exemplary bool      `json:"exemplary"`
	Enabled   bool      `json:"enabled"`
	Color     ItemColor `json:"color"`
	Size      ItemSize  `json:"size"`
}

// ItemColor is a color of Item.
type ItemColor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RGB  string `json:"rgb"`
}

// ItemSize is a size of Item.
type ItemSize struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetItems gets all item list.
func (c *Client) GetItems(ctx context.Context) ([]Item, error) {
	resp, err := c.get(ctx, "/items", nil)
	if err != nil {
		return nil, err
	}

	// TODO: status chack and error handling
	var items Items
	if err := decodeJSON(resp, &items); err != nil {
		return nil, err
	}

	return items.Items, nil
}
