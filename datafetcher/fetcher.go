package datafetcher

import (
	"io"
	"net/http"
	"sol/floortracker/constants"
)

func GetStats(symbol string) string {
	url := constants.StatsReq(symbol)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body)
}
