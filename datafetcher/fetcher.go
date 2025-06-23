package datafetcher

import (
	"io"
	"net/http"
	"sol/floortracker/constants"
)

func GetStats(symbol string, ch chan []byte) error {
	url := constants.StatsReq(symbol)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	ch <- body

	return nil
}
