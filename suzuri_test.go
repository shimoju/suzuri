package suzuri

import (
	"net/http"
	"net/http/httptest"
)

func stubServer() (*http.ServeMux, *httptest.Server) {
	stub := http.NewServeMux()
	server := httptest.NewServer(stub)
	return stub, server
}
