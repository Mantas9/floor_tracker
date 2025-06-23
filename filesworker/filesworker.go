package filesworker

import (
	"encoding/json"
	"os"
	"sol/floortracker/constants"
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

func (s *Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(s)
}

func WriteJSON(data []byte) error {
	err := os.WriteFile("./stats.json", data, 0644) // Write string data to file

	return err
}
