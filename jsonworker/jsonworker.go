package jsonworker

import (
	"encoding/json"
	"sol/floortracker/constants"
)

// / Result structs
type Stats struct {
	Symbol      string
	FloorPrice  float64
	ListedCount float64
}

func (s *Stats) UnmarshalJSON(data []byte) error {
	var statsData map[string]interface{} // Placeholder data

	// Unmarshal JSON
	err := json.Unmarshal(data, &statsData)
	// Error check
	if err != nil {
		return err
	}

	s.Symbol = statsData["symbol"].(string)
	s.FloorPrice = statsData["floorPrice"].(float64) * constants.LamportsToSol
	s.ListedCount = statsData["listedCount"].(float64)

	return nil
}
