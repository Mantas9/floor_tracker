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
	if len(args) == 0 {
		fmt.Println("Usage: program [-t | -j | -c] collection1 collection2 ...")
		os.Exit(1)
	}

	// Default to terminal
	export := Terminal

	// Check for output mode
	if args[0][0] == '-' {
		switch args[0][1] {
		case 't', 'T':
			export = Terminal
		case 'j', 'J':
			export = JSON
		case 'c', 'C':
			export = CSV
		default:
			fmt.Println("Unknown flag:", args[0])
			fmt.Println("Valid flags: -t (Terminal), -j (JSON), -c (CSV)")
			os.Exit(1)
		}
		args = args[1:] // Remove flag from args
	}

	if len(args) == 0 {
		fmt.Println("Error: No collections specified.")
		os.Exit(1)
	}

	var wg sync.WaitGroup // Create waitgroup to let code finish before exiting

	// Fetch all requested data
	dataChan := make(chan filesworker.Stats) // Channel for concurrent communication

	// Iterate through arguments
	for _, symbol := range args {
		wg.Add(1)

		// Get data and send to channel
		go func(sym string) {
			defer wg.Done()

			stat, err := getStats(sym)

			if err != nil {
				fmt.Printf("Error in fetching stats:\n%s", err)
				os.Exit(1)
			}

			// Push to channel when done
			dataChan <- stat
		}(symbol)
	}

	// Wait for workflow to finish
	go func() {
		wg.Wait()       // Waitgroup block
		close(dataChan) // Close channel when everything is done
	}()

	// Iterate through channel data and push everything to array
	var data []filesworker.Stats
	for stat := range dataChan {
		data = append(data, stat) // Append
	}

	// Onde data list is filled, time to export accordingly
	switch export {
	case Terminal:
		exportTerminal(data)
	case JSON:
		filesworker.WriteJSON(data)
	case CSV:
		filesworker.WriteCSV(data)
	}

	wg.Wait()
}

func exportTerminal(data []filesworker.Stats) {
	// Iterate through data
	for _, stats := range data {
		// Print to terminal
		fmt.Println("==============================")
		fmt.Println("Collection Name:", stats.Symbol)
		fmt.Println("Floor Price:", stats.FloorPrice, "SOL")
		fmt.Println("Listed Count:", stats.ListedCount)
		fmt.Println("==============================")
		fmt.Println() // Space
	}
}

// Returns NFT stats in the corresponding struct format
func getStats(symbol string) (filesworker.Stats, error) {

	// Make a channel for communicating with HTTP data
	ch := make(chan []byte, 1)

	// Get online data
	if err := datafetcher.GetStats(symbol, ch); err != nil {
		return filesworker.Stats{}, err
	}

	// Block until datafetcher finishes and dump data to rawData variable
	rawData := <-ch

	var stats filesworker.Stats // Result variable

	if err := stats.UnmarshalJSON(rawData); err != nil {
		return filesworker.Stats{}, err
	}

	return stats, nil
}
