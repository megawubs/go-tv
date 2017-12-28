package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hdhauk/tv/tv"
)

// Send will send a command to a server listening on localhost:8000
func Send(c tv.Command) error {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(c); err != nil {
		return fmt.Errorf("failed to encode json: %v", err)
	}

	if _, err := http.Post("http://localhost:8000/", "application/json; charset=utf-8", b); err != nil {
		return fmt.Errorf("failed to post to server: %v", err)
	}

	return nil
}
