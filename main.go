package main

import (
	"fmt"
	"flag"
	"github.com/mstch/gotrack/provider"
)

func main() {
	// Define command-line flags
	tracker := flag.String("p", "", "Tracker")
	packageID := flag.String("i", "", "Package ID")

	// Parse command-line arguments
	flag.Parse()

	// Check if the required flags are provided
	if *tracker == "" || *packageID == "" {
		fmt.Println("Usage: gotracker -p [TRACKER] -i [PACKAGE_ID]")
		return
	}

	request := provider.NewRequest(*tracker, *packageID)

	fmt.Println(request)
}
