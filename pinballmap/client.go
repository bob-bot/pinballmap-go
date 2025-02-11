// Package pinballmap provides an SDK for interacting with the Pinball Map API.
package pinballmap

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://pinballmap.com/api/v1"
)

// Client represents a Pinball Map API client.
type Client struct {
	APIKey string
}

// NewClient initializes a new API client.
func (c *Client) get(endpoint string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(fmt.Sprintf("%s/%s", baseURL, endpoint))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	// Decode the response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}