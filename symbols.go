package gemini

type (
	SymbolsResponse []string

	SymbolDetailsResponse struct {
		Symbol         string  `json:"symbol"`
		BaseCurrency   string  `json:"base_currency"`
		QuoteCurrency  string  `json:"quote_currency"`
		TickSize       float64 `json:"tick_size"`
		QuoteIncrement float64 `json:"quote_increment"`
		MinOrderSize   string  `json:"min_order_size"`
		Status         string  `json:"status"`
	}
)

// GetSymbols Returns available symbols at Gemini
func (c *Client) GetSymbols() (resp SymbolsResponse, err error) {
	err = c.Call("GET", "/symbols", nil, &resp)
	return
}

// GetSymbolDetails Returns detailed information about symbol such as decimal sizes, min order size, currency information.
func (c *Client) GetSymbolDetails(symbol string) (resp SymbolDetailsResponse, err error) {
	err = c.Call("GET", "/symbols/details/"+symbol, nil, &resp)
	return
}
