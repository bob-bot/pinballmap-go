package pinballmap

import (
	"encoding/json"
	"fmt"
)

// GetMachines fetches all available pinball machines.
func (c *Client) GetMachines() ([]Machine, error) {
	endpoint := "machines.json"

	data, err := c.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Machines []Machine `json:"machines"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parse machines response: %w", err)
	}

	return result.Machines, nil
}