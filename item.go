package suzuri

import (
	"context"
)

type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Angles       []string  `json:"angles"`
	HumanizeName string    `json:"humanizeName"`
	Variants     []Variant `json:"variants"`
}

type Variant struct {
	ID        int   `json:"id"`
	Price     int   `json:"price"`
	Exemplary bool  `json:"exemplary"`
	Enabled   bool  `json:"enabled"`
	Color     Color `json:"color"`
	Size      Size  `json:"size"`
}

type Color struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	RGB  string `json:"rgb"`
}

type Size struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) GetItems(ctx context.Context) ([]Item, error) {
	resp, err := c.get(ctx, "/items", nil)
	if err != nil {
		return nil, err
	}

	var items Items
	if err := decodeJSON(resp, &items); err != nil {
		return nil, err
	}

	return items.Items, nil
}
