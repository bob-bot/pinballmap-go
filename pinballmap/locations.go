package pinballmap

import (
	"encoding/json"
	"fmt"
)

// GetLocations fetches pinball locations based on a city.
func (c *Client) GetLocations(city string) ([]Location, error) {
	endpoint := "locations.json"
	params := map[string]string{"by_city": city}

	data, err := c.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Locations []Location `json:"locations"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse locations response: %w", err)
	}

	return result.Locations, nil
}