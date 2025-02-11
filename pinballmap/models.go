package pinballmap

// Location represents a pinball location.
type Location struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	City     string  `json:"city"`
	RegionID int     `json:"region_id"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

// Machine represents a pinball machine.
type Machine struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Region represents a pinball region. (This stays here!)
type Region struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}