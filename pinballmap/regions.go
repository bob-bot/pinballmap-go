package pinballmap

import (
	"encoding/json"
	"fmt"
)

// GetRegions fetches all available pinball regions.
func (c *Client) GetRegions() ([]Region, error) {
	endpoint := "regions.json"

	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Regions []Region `json:"regions"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse regions response: %w", err)
	}

	return result.Regions, nil
}