package suzuri

import (
	"log"
	"net/http"
	"net/url"
)

// Client is a SUZURI client for making SUZURI API requests.
type Client struct {
	URL        *url.URL
	HTTPClient *http.Client

	Token string

	Logger *log.Logger
}
