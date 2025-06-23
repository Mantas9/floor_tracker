package filesworker

import (
	"encoding/json"
	"os"
	"sol/floortracker/constants"

	"github.com/gocarina/gocsv"
)

// / Result structs
type Stats struct {
	Symbol      string  `csv:"symbol" json:"symbol"`
	FloorPrice  float64 `csv:"floorPrice" json:"floorPrice"`
	ListedCount float64 `csv:"listedCount" json:"listedCount"`
}

func (s *Stats) UnmarshalJSON(data []byte) error {
	var statsData map[string]interface{} // Placeholder data

	// Unmarshal JSON
	err := json.Unmarshal(data, &statsData)
	// Error check
	if err != nil {
		return err
	}

	// Assign JSON values
	s.Symbol = statsData["symbol"].(string)
	s.FloorPrice = statsData["floorPrice"].(float64) * constants.LamportsToSol
	s.ListedCount = statsData["listedCount"].(float64)

	return nil
}

func WriteJSON(data []Stats) error {
	// Marshal data to JSON
	jsonData, err := json.Marshal(data)

	if err != nil { // Error check
		return err
	}

	err = os.WriteFile("./stats.json", jsonData, 0644) // Write string data to file

	return err
}

func WriteCSV(data []Stats) error {
	// Create result CSV file
	file, err := os.Create("stats.csv")
	if err != nil {
		return err
	}

	// Marshal data struct list to CSV and write to file
	if err := gocsv.MarshalFile(&data, file); err != nil {
		return err
	}

	return nil
}
