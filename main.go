package main

import (
	"fmt"
	"log"

	"github.com/bob-bot/pinballmap-go/pinballmap"
)

func main() {
	// Initialize the Pinball Map API client (API key is optional for public queries)
	client := pinballmap.NewClient("")

	// Test fetching locations
	fmt.Println("Fetching locations in New York...")
	locations, err := client.GetLocations("New York")
	if err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}

	// Print first few locations
	for i, loc := range locations {
		if i >= 5 {
			break
		}
		fmt.Printf("Location %d: %s (%s)\n", loc.ID, loc.Name, loc.City)
	}

	// Test fetching machines
	fmt.Println("\nFetching available pinball machines...")
	machines, err := client.GetMachines()
	if err != nil {
		log.Fatalf("Error fetching machines: %v", err)
	}

	// Print first few machines
	for i, mach := range machines {
		if i >= 5 {
			break
		}
		fmt.Printf("Machine %d: %s\n", mach.ID, mach.Name)
	}

	// Test fetching regions
	fmt.Println("\nFetching available regions...")
	regions, err := client.GetRegions()
	if err != nil {
		log.Fatalf("Error fetching regions: %v", err)
	}

	// Print first few regions
	for i, reg := range regions {
		if i >= 5 {
			break
		}
		fmt.Printf("Region %d: %s (%s)\n", reg.ID, reg.Name, reg.Slug)
	}
}