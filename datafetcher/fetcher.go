package datafetcher

import (
	"io"
	"net/http"
	"sol/floortracker/constants"
)

func GetStats(symbol string, ch chan []byte) {
	url := constants.StatsReq(symbol)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	ch <- body
}
