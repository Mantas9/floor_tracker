package main

import (
	"fmt"
	"sol/floortracker/datafetcher"
)

func main() {
	fmt.Println(datafetcher.GetStats("degods"))
}
