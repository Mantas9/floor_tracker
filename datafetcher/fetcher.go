package datafetcher

import (
	"net/http"
	"sol/floortracker/constants"
)

func getFloorPrice(symbol string) string {
	url := constants.FloorReq(symbol)

	req, _ := http.NewRequest("GET", url, nil)
}
