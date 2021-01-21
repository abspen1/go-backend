package alp

import (
	"os"
	"strconv"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

type Ticker struct {
	Stock string
}

type alpacaClientContainer struct {
	client  *alpaca.Client
	stock   string
	amtBars int
}

var alpacaClient alpacaClientContainer

func initialize(ticker string) {
	os.Setenv(common.EnvApiKeyID, os.Getenv("ALP-KEY"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("ALP-SECRET"))

	// fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	alpaca.SetBaseUrl(os.Getenv("https://paper-api.alpaca.markets"))
	alpacaClient = alpacaClientContainer{
		alpaca.NewClient(common.Credentials()),
		"stock",
		10,
	}
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func (alp alpacaClientContainer) getCurrPrice() string {
	bars, _ := alp.client.GetSymbolBars(alpacaClient.stock, alpaca.ListBarParams{Timeframe: "minute", Limit: &alpacaClient.amtBars})
	currPrice := float64(bars[len(bars)-1].Close)
	price := FloatToString(currPrice)
	return price
}

func GetCurrentPrice(ticker string) string {
	initialize(ticker)
	return alpacaClient.getCurrPrice()
}
