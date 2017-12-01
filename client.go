package suzuri

import (
	"context"
	"encoding/json"
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

// SetBaseURL changes the base URL for API requests.
func (c *Client) SetBaseURL(urlStr string) error {
	baseURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	c.baseURL = baseURL

	return nil
}

func (c *Client) get(ctx context.Context, endpoint string, params url.Values) (*http.Response, error) {
	req, _ := c.newRequest(ctx, "GET", endpoint, params, nil)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) newRequest(ctx context.Context, method, endpoint string, query url.Values, body io.Reader) (*http.Request, error) {
	reqURL := *c.baseURL
	reqURL.Path = path.Join(c.baseURL.Path, endpoint)
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest(method, reqURL.String(), body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", userAgent)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func decodeJSON(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(result)
}
