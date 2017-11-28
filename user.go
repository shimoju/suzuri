package suzuri

import (
	"context"
	"strconv"
)

// User is a SUZURI user account.
type User struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	DisplayName string         `json:"displayName"`
	AvatarURL   string         `json:"avatarUrl"`
	Identities  []UserIdentity `json:"identities"`
	Profile     UserProfile    `json:"profile"`
}

// UserIdentity is information about social account connected with a User.
type UserIdentity struct {
	ID       int    `json:"id"`
	Provider string `json:"provider"`
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
	URL      string `json:"url"`
}

// UserProfile is a profile of User.
type UserProfile struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Body      string `json:"body"`
	HeaderURL string `json:"headerUrl"`
}

// userResponse is a response data structure when API returns a User.
type userResponse struct {
	User User `json:"user"`
}

// GetUser gets details about an existing user.
func (c *Client) GetUser(ctx context.Context, userID int) (*User, error) {
	endpoint := "/users/" + strconv.Itoa(userID)
	resp, err := c.get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}

	// TODO: status chack and error handling
	var userResp userResponse
	if err := decodeJSON(resp, &userResp); err != nil {
		return nil, err
	}

	return &userResp.User, nil
}
