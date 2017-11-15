package suzuri

import (
	"net/http"
	"net/url"
)

const baseURL = "https://suzuri.jp"

// Client is a SUZURI client for making SUZURI API requests.
type Client struct {
	url        *url.URL
	httpClient *http.Client

	token string
}

// NewClient returns a new Client.
func NewClient(token string) *Client {
	parsedURL, _ := url.ParseRequestURI(baseURL)

	return &Client{
		url:        parsedURL,
		httpClient: http.DefaultClient,
		token:      token,
	}
}
