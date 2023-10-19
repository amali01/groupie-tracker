package apiprocess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ParseAPI sends a GET request to the specified URL and unmarshals the response
// body into the provided interface.
func ParseAPI(url string, d interface{}) error {
	// Send GET request to URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %s", err)
	}

	// Unmarshal response body into provided interface
	if err := json.Unmarshal(body, d); err != nil {
		return fmt.Errorf("error unmarshaling response: %s", err)
	}
	return nil
}
