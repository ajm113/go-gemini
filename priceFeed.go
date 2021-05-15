package gemini

type (
	CurrencyPrice struct {
		Pair             string `json:"pair"`
		Price            string `json:"price"`
		PercentChange24h string `json:"percentChange24h"`
	}

	PriceFeedResponse []CurrencyPrice
)

// GetPriceFeed Returns current market prices from Gemini
// @see https://docs.gemini.com/rest-api/#price-feed
func (c *Client) GetPriceFeed() (resp PriceFeedResponse, err error) {
	err = c.Call("GET", "/pricefeed", nil, &resp)
	return
}
