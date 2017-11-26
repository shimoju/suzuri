package suzuri

import "context"

// UserRoot is the root element that wraps a User.
type UserRoot struct {
	User User `json:"user"`
}

// User is a SUZURI user account.
type User struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	DisplayName string     `json:"displayName"`
	AvatarURL   string     `json:"avatarUrl"`
	Identities  []Identity `json:"identities"`
	Profile     Profile    `json:"profile"`
}

// Identity is information about social account connected with a User.
type Identity struct {
	ID       int    `json:"id"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
	URL      string `json:"url"`
}

// Profile is a profile of User.
type Profile struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Body      string `json:"body"`
	HeaderURL string `json:"headerUrl"`
}

// GetUser gets details about an existing user.
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	endpoint := "/users/" + userID
	resp, err := c.get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// TODO: status chack and error handling
	var user UserRoot
	if err := decodeJSON(resp, &user); err != nil {
		return nil, err
	}

	return &user.User, nil
}
