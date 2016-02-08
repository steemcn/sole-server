package utils

import (
	"encoding/json"

	"github.com/freeusd/solebtc/Godeps/_workspace/src/github.com/parnurzeal/gorequest"
)

const blockchainTickerURL = "https://blockchain.info/ticker"

// BitcoinPrice returns the lastest bitcoin price in USD
func BitcoinPrice() (_ float64, err error) {
	_, body, _ := gorequest.New().Get(blockchainTickerURL).EndBytes()
	return bitcoinPriceWithByteFromBlockchain(body)
}

func bitcoinPriceWithByteFromBlockchain(data []byte) (float64, error) {
	m := map[string]struct { // use pointer so it fails fast
		Last float64 `json:"last"`
	}{}

	err := json.Unmarshal(data, &m)
	if err != nil {
		return 0, err
	}

	return m["USD"].Last, nil
}
