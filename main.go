package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sol/floortracker/datafetcher"
	"sol/floortracker/filesworker"
	"sync"

	"github.com/gocarina/gocsv"
)

const ( // Export As variables
	Terminal = iota
	JSON
	CSV
)

func main() {
	// Get OS arguments (skip run program variable)
	args := os.Args[1:]

	// Export variable
	export := -1

	// Get export option if specified. Else, default to terminal output
	if args[0][0] == '-' {
		params := args[0] // Parameters string

		for i := 1; i < len(params); i++ { // Go through each letter argument
			if export == -1 && (params[i] == 't' || params[i] == 'T') { // Terminal output
				export = Terminal
			}
			if export == -1 && (params[i] == 'j' || params[i] == 'J') { // JSON output
				export = JSON
			}
			if export == -1 && (params[i] == 'c' || params[i] == 'C') { // CSV output
				export = CSV
			}
		}

		// Truncate parameters
		args = args[1:]
	} else { // Default to terminal output
		export = Terminal
	}

	var wg sync.WaitGroup // Create waitgroup to let code finish before exiting

	// Fetch all requested data
	var data []filesworker.Stats       // Empty list of stats
	ch := make(chan filesworker.Stats) // Channel for concurrent communication

	for i := 0; i < len(args); i++ {
		// Call getStats
		getStats(args[i], ch)

		// Append new stats to list
		data = append(data, <-ch)
	}

	fmt.Println(data[0])

	// Onde data list is filled, time to export accordingly
	switch export {
	case Terminal: // Prints each collection's stats to terminal

		// Iterate through data
		for i := 0; i < len(data); i++ {
			stats := data[i] // Idividual stats variable of single collection

			// Print to terminal
			fmt.Println("==============================")
			fmt.Println("Collection Name:", stats.Symbol)
			fmt.Println("Floor Price:", stats.FloorPrice, "SOL")
			fmt.Println("Listed Count:", stats.ListedCount)
			fmt.Println("==============================")
			fmt.Println() // Space
		}

	case JSON:

		// Marshal data to JSON
		jsonData, err := json.Marshal(data)

		if err != nil { // Error check
			panic(err)
		}

		// Write JSON to file
		wg.Add(1)
		go func() {
			defer wg.Done()

			filesworker.WriteJSON(jsonData)
		}()

	case CSV:

		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create result CSV file
			file, err := os.Create("stats.csv")
			if err != nil {
				panic(err)
			}

			// Marshal data struct list to CSV and write to file
			if err := gocsv.MarshalFile(&data, file); err != nil {
				panic(err)
			}
		}()

	}

	wg.Wait()
}

// Returns NFT stats in the corresponding struct format
func getStats(symbol string, rc chan filesworker.Stats) {
	// Make a channel for communicating with HTTP data
	ch := make(chan []byte)

	// Get online data
	datafetcher.GetStats(symbol, ch)

	// Block until datafetcher finishes and dump data to rawData variable
	rawData := <-ch

	var stats filesworker.Stats // Result variable

	err := stats.UnmarshalJSON(rawData)
	if err != nil {
		panic(err)
	}

	rc <- stats
}
