package suzuri

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
)

const suzuriURL = "https://suzuri.jp"

var userAgent = fmt.Sprintf("SuzuriGo/%s (%s)", version, runtime.Version())

// Client is a SUZURI client for making SUZURI API requests.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	token string
}

// NewClient returns a new Client.
func NewClient(token string) *Client {
	baseURL, _ := url.ParseRequestURI(suzuriURL)

	return &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		token:      token,
	}
}
