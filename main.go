package main

import (
	"fmt"
	"sol/floortracker/datafetcher"
	"sol/floortracker/jsonworker"
)

func main() {
	getStats("degods")
}

func getStats(symbol string) {
	ch := make(chan []byte)

	go datafetcher.GetStats(symbol, ch)

	rawData := <-ch

	var stats jsonworker.Stats

	stats.UnmarshalJSON(rawData)

	fmt.Println(stats.Symbol, stats.FloorPrice, stats.ListedCount)
}
