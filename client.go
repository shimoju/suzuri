package suzuri

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

const suzuriAPI = "https://suzuri.jp/api/v1"

var userAgent = fmt.Sprintf("SuzuriGo/%s (%s)", version, runtime.Version())

// Client is a SUZURI client for making SUZURI API requests.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	token string
}

// NewClient returns a new Client.
func NewClient(token string) *Client {
	baseURL, _ := url.ParseRequestURI(suzuriAPI)

	return &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
		token:      token,
	}
}

func (c *Client) newRequest(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	reqURL := *c.baseURL
	reqURL.Path = path.Join(c.baseURL.Path, endpoint)

	req, err := http.NewRequest(method, reqURL.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	return req, nil
}
