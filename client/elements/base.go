// Copyright 2016 Palm Stone Games, Inc. All rights reserved.

package elements

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func DoHTTPJSON(method string, url string, data interface{}) (*http.Response, error) {
	// Encode message
	encodedData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling JSON to %v on %v: %v", method, url, err)
	}

	// Create HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(encodedData))
	if err != nil {
		return nil, fmt.Errorf("Error while creating request: %v", err)
	}

	// Do HTTP call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while submitting DOWN request: %v", err)
	}

	// Check status
	if resp.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("Got http status %v during DOWN request instead of expected 200 OK", resp.Status)
	}

	return resp, nil
}
