package pinballmap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetLocations fetches pinball locations based on a city.
func (c *Client) GetLocations(city string) ([]Location, error) {
	url := fmt.Sprintf("%s/locations.json?by_city=%s", c.BaseURL, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch locations, status code: %d", resp.StatusCode)
	}

	var result struct {
		Locations []Location `json:"locations"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Locations, nil
}