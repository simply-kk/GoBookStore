package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads and unmarshals the request body into the provided struct
func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, x)
	if err != nil {
		return
	}
}
