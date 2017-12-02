package suzuri

import (
	"encoding/json"
	"net/http"
)

func decodeJSON(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(result)
}
