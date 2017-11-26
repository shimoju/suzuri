package suzuri

import (
	"context"
	"net/http"
	"net/http/httptest"
)

var (
	stub   *http.ServeMux
	server *httptest.Server
	ctx    context.Context
	client *Client
)

func setup() {
	stub, server = stubServer()
	client = NewClient("accesstoken")
	client.SetBaseURL(server.URL)
	ctx = context.Background()
}

func teardown() {
	server.Close()
}

func stubServer() (*http.ServeMux, *httptest.Server) {
	stub := http.NewServeMux()
	server := httptest.NewServer(stub)
	return stub, server
}
