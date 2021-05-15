package gemini

import "encoding/json"

type (
	OrderStatusRequestPayload struct {
		BaseRequest
		OrderID       int  `json:"order_id"`
		IncludeTrades bool `json:"include_trades"`
	}

	OrderStatusTrade struct {
		Aggressor     bool   `json:"aggressor"`
		Amount        bool   `json:"amount"`
		Exchange      string `json:"exchange"`
		FeeAmount     string `json:"fee_amount"`
		FeeCurrency   string `json:"fee_currency"`
		IsAuctionFill bool   `json:"is_auction_fill"`
		OrderID       string `json:"order_id"`
		Price         string `json:"price"`
		TID           int    `json:"tid"`
		Timestamp     int64  `json:"timestamp"`
		TimestampMS   int64  `json:"timestampms"`
		Type          string `json:"type"`
	}

	OrderStatusResponseWithTrades struct {
		OrderStatusResponse
		Trades []OrderStatusTrade `json:"trades"`
	}
)

// OrderStatus Returns status of order.
// WARNING: Gemini recommends using WebSocket Order Events API to receive order status changes!
// WARNING: Under the terms of the Gemini API Agreement, polling this endpoint may be subject to rate limiting.
// @see https://docs.gemini.com/rest-api/#order-status
func (c *Client) GetOrderStatus(orderID int) (resp OrderStatusResponse, err error) {

	order := OrderStatusRequestPayload{
		OrderID: orderID,
	}

	order.Nonce = c.createNonce()
	order.Request = "/v1/order/status"

	bodyBytes, err := json.Marshal(order)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/status", bodyBytes, &resp)

	return
}

// GetOrderStatusWithTrades Returns status of order including trades.
// WARNING: Gemini recommends using WebSocket Order Events API to receive order status changes!
// WARNING: Under the terms of the Gemini API Agreement, polling this endpoint may be subject to rate limiting.
// @see https://docs.gemini.com/rest-api/#order-status
func (c *Client) GetOrderStatusWithTrades(orderID int) (resp OrderStatusResponseWithTrades, err error) {

	order := OrderStatusRequestPayload{
		OrderID:       orderID,
		IncludeTrades: true,
	}

	order.Nonce = c.createNonce()
	order.Request = "/v1/order/status"

	bodyBytes, err := json.Marshal(order)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/status", bodyBytes, &resp)

	return
}

//GetActiveOrders Returns ALL active orders in Gemini.
//@see https://docs.gemini.com/rest-api/#get-active-orders
func (c *Client) GetActiveOrders() (resp []OrderStatusResponse, err error) {
	body := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/orders",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/orders", bodyBytes, &resp)

	return
}

type (
	GetPastTradesOptions struct {
		Symbol      string
		LimitTrades int
		Timestamp   int64
		Account     string
	}

	GetPastTradesPayload struct {
		BaseRequest
		Symbol      string `json:"symbol"`
		LimitTrades int    `json:"limit_trades,omitempty"`
		Timestamp   int64  `json:"timestamp,omitempty"`
		Account     string `json:"account,omitempty"`
	}
)

//GetPastTrades Returns all past trades made by user with supplied symbol
//@see https://docs.gemini.com/rest-api/#get-past-trades
func (c *Client) GetPastTrades(options GetPastTradesOptions) (resp []OrderStatusTrade, err error) {
	body := GetPastTradesPayload{
		Symbol:      options.Symbol,
		LimitTrades: options.LimitTrades,
		Timestamp:   options.Timestamp,
		Account:     options.Account,
	}

	body.Nonce = c.createNonce()
	body.Request = "/v1/mytrades"

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/mytrades", bodyBytes, &resp)

	return
}
