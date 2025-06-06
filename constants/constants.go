package constants

const LamportsToSol = 1_000_000_000

const apiURL = "https://api-mainnet.magiceden.dev/v2/"

func FloorReq(symbol string) string {
	return apiURL + "collections/" + symbol + "/stats"
}
