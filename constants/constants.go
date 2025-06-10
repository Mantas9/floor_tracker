package constants

// / Misc. Constants
const LamportsToSol = 0.000_000_001

const apiURL = "https://api-mainnet.magiceden.dev/v2/"

// / Requests
func StatsReq(symbol string) string {
	return apiURL + "collections/" + symbol + "/stats"
}
