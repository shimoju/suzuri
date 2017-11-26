package suzuri

import "context"

// User is a SUZURI user account.
type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl"`
	Identities  []struct {
		ID       int    `json:"id"`
		Provider string `json:"provider"`
		UID      string `json:"uid"`
		Nickname string `json:"nickname"`
		URL      string `json:"url"`
	} `json:"identities"`
	Profile struct {
		ID        int    `json:"id"`
		URL       string `json:"url"`
		Body      string `json:"body"`
		HeaderURL string `json:"headerUrl"`
	} `json:"profile"`
}

// GetUser gets details about an existing user.
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	user := &User{ID: 1}

	return user, nil
}
