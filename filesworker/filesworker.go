package filesworker

import (
	"encoding/json"
	"os"
	"sol/floortracker/constants"
)

// / Result structs
type Stats struct {
	Symbol      string
	FloorPrice  float64
	ListedCount float64

	RawData []byte
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
	s.RawData = data // original raw data

	return nil
}

func (s *Stats) WriteToJSON() error {
	// Write to JSON
	err := os.WriteFile("./data.json", s.RawData, 0644)

	return err
}
