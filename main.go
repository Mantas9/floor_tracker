package main

import (
	"fmt"
	"os"
	"sol/floortracker/datafetcher"
	"sol/floortracker/filesworker"
	"sync"
)

const ( // Export As variables
	Terminal = iota
	JSON
	CSV
)

func main() {
	// Get OS arguments (skip run program variable)
	args := os.Args[1:]

	// Default export variable
	export := Terminal

	var wg sync.WaitGroup // Create waitgroup to let code finish before exiting

	// Cycle through the designated arguments and call getStats
	for i := 0; i < len(args); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exportStats(args[i], export)
		}
	}

	wg.Wait()
}

// Gets stats and exports them according to user's choice
func exportStats(arg string, exportAs int) {
	// Create channel to block for getStats
	ch := make(chan filesworker.Stats)

	// Get stats from argument
	go getStats(arg, ch)

	// Channel data to variable
	stats := <-ch

	// ExportAs logic
	switch exportAs {
	case Terminal: // Simply print the values to terminal
		fmt.Println("==============================")
		fmt.Println("Collection Name:", stats.Symbol)
		fmt.Println("Floor Price:", stats.FloorPrice)
		fmt.Println("Listed Count:", stats.ListedCount)
		fmt.Println("==============================")

	case JSON: // Write RAW data to a JSON file
		err := stats.WriteToJSON()

		// Panic if something goes wrong
		if err != nil {
			panic(err)
		}
	}
}

// Returns NFT stats in the corresponding struct (declared in jsonworker module)
func getStats(symbol string, rc chan filesworker.Stats) {
	// Make a channel for communicating with HTTP data
	ch := make(chan []byte)

	// Run goroutine to get online data
	go datafetcher.GetStats(symbol, ch)

	// Block until datafetcher finishes and dump data to rawData variable
	rawData := <-ch

	var stats filesworker.Stats // Result variable

	// Fill with designated values
	stats.UnmarshalJSON(rawData)

	rc <- stats
}
