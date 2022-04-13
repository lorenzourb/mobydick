package utils 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"math/big"
	"strconv"
	"fmt"
	"github.com/Jeffail/gabs"
)

var (
 	MainnetEndpoint = os.Getenv("MAINNET_ENDPOINT")
	PostgresEndpoint = os.Getenv("POSTGRES_ENDPOINT")
	CoinMarketCapURL = os.Getenv("COINMARKETCAP_URL")
)

const (
	TxAmountThreshold = 100000000000
	SixDigitsThreshold = 1000000
	EighteenDigitsThreshold = 1000000000000000000
)

type Token struct {
	Name string
	Address string
	Decimals int
	ID string
}

var Tokens = []Token{
	Token {
		Name: "Tether",
		Address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
		Decimals: 6,
		ID: "825",
	},
	Token {
		Name: "BNB",
		Address: "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
		Decimals: 18,
		ID: "1839",
	}, 
	Token {
		Name: "WLUNA",
		Address: "0xd2877702675e6cEb975b4A1dFf9fb7BAF4C91ea9",
		Decimals: 18,
		ID: "4172",
	},
	Token {
		Name: "USDC",
		Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		Decimals: 6,
		ID: "3408",
	},
	Token {
		Name: "BUSD",
		Address: "0x4Fabb145d64652a948d72533023f6E7A623C7C53",
		Decimals: 18,
		ID: "4687",
	},
}

func HexToBigInt(s string) *big.Int{
	numberStr := strings.Replace(s, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	i := new(big.Int)
	i.SetString(numberStr, 16)
	return i
}

func HexToInt(s string) int64{
	numberStr := strings.Replace(s, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	output, err := strconv.ParseInt(numberStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return output
}

type Resp struct {
	status interface{} `json:"status"`
	Data  json.RawMessage `json:"data"`
}

func FetchAllTokenPrices() *gabs.Container{
	var tokenIds []string
	for _,t := range Tokens {
		tokenIds = append(tokenIds, t.ID)
	}

	resp, err := http.Get(CoinMarketCapURL + strings.Join(tokenIds[:], ","))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) 

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	return jsonParsed
}

func FilterTokenPrice(container *gabs.Container, tokenID string) string{
	return container.Path("data." + tokenID + ".quote.USD.price").String()
}