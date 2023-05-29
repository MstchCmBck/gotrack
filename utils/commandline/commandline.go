package commandline

import (
	"flag"
	"fmt"
)

type Args struct {
	tracker string
	packageId string
}

func (a *Args) Parse(args []string) {
	// Define command-line flags
	tracker := flag.String("p", "", "Tracker")
	packageId := flag.String("i", "", "Package ID")

	// Parse command-line arguments
	flag.CommandLine.Parse(args)

	// Fill Args struct
	a.tracker = *tracker
	a.packageId = *packageId

	// Check if the required flags are provided
	if a.tracker == "" || a.packageId == "" {
		fmt.Println("Usage: gotracker -p [TRACKER] -i [PACKAGE_ID]")
		return
	}
}

func (a Args) GetTracker() string {
	return a.tracker
}

func (a Args) GetPackageId() string {
	return a.packageId
}
