package alp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
)

// Ticker struct holds stock
type Ticker struct {
	Stock string
}

type alpacaClientContainer struct {
	client  *alpaca.Client
	stock   string
	amtBars int
}

var alpacaClient alpacaClientContainer

// Get func just displays simple text at the endpoint
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<body style="text-align:center;">
	<h1>Current stock price endpoint written in Go!<h1>
	<img src="https://mms.businesswire.com/media/20181023005433/en/686001/5/Alpaca_Logo_yellow.jpg" alt="Alpaca Logo">
	</body>`)
}

// PostStockPrice func is endpoint Post request
func PostStockPrice(w http.ResponseWriter, r *http.Request) {
	var body []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ = ioutil.ReadAll(r.Body)
	}
	var stock Ticker
	_ = json.Unmarshal(body, &stock)

	fmt.Println(stock.Stock)
	resp := GetCurrentPrice(stock.Stock)
	fmt.Fprintf(w, resp)
}

func initialize(ticker string) {
	os.Setenv(common.EnvApiKeyID, os.Getenv("APCA_API_KEY_ID"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("APCA_API_SECRET_KEY"))

	fmt.Printf("Running w/ credentials [%v %v]\n", common.Credentials().ID, common.Credentials().Secret)

	alpaca.SetBaseUrl(os.Getenv("APCA_API_BASE_URL"))
	alpacaClient = alpacaClientContainer{
		alpaca.NewClient(common.Credentials()),
		ticker,
		10,
	}
}

func floatToString(num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(num, 'f', 6, 64)
}

func (alp alpacaClientContainer) getCurrPrice() string {
	bars, err := alp.client.GetSymbolBars(alpacaClient.stock, alpaca.ListBarParams{Timeframe: "minute", Limit: &alpacaClient.amtBars})
	if err != nil {
		fmt.Println(err)
		return "500"
	}
	if len(bars) <= 0 {
		return "404"
	}
	currPrice := float64(bars[len(bars)-1].Close)
	price := floatToString(currPrice)
	return price
}

// GetCurrentPrice func returns the current price of given ticker
func GetCurrentPrice(ticker string) string {
	initialize(ticker)
	return alpacaClient.getCurrPrice()
}
