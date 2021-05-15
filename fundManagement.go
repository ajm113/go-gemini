package gemini

import "encoding/json"

type (
	AvailableBalance struct {
		Type                   string `json:"type"`
		Currency               string `json:"currency"`
		Amount                 string `json:"amount"`
		Available              string `json:"available"`
		AvailableForWithdrawal string `json:"availableForWithdrawal"`
	}

	AvailableBalanceResponse []AvailableBalance
)

//GetAvailableBalances Returns all balances in the exchange and available amounts to trade with.
//@see https://docs.gemini.com/rest-api/#get-available-balances
func (c *Client) GetAvailableBalances() (resp AvailableBalanceResponse, err error) {
	body := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/balances",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/balances", bodyBytes, &resp)

	return
}
