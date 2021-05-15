package gemini

import "encoding/json"

type (
	NewOrderPayload struct {
		BaseRequest
		Symbol        string   `json:"symbol"`
		Amount        string   `json:"amount"`
		Price         string   `json:"price"`
		ClientOrderID string   `json:"client_order_id,omitempty"`
		StopPrice     string   `json:"stop_price,omitempty"`
		Side          string   `json:"side"`
		Type          string   `json:"type"`
		Options       []string `json:"options,omitempty"`
	}

	OrderStatusResponse struct {
		OrderID           string   `json:"order_id"`
		ID                string   `json:"id"`
		Symbol            string   `json:"symbol"`
		Exchange          string   `json:"exchange"`
		AVGExecutionPrice string   `json:"avg_execution_price"`
		Side              string   `json:"side"`
		Type              string   `json:"type"`
		Timestamp         string   `json:"timestamp"`
		Timestampms       int64    `json:"timestampms"`
		IsLive            bool     `json:"is_live"`
		IsCancelled       bool     `json:"is_cancelled"`
		IsHidden          bool     `json:"is_hidden"`
		WasForced         bool     `json:"was_forced"`
		ExecutedAmount    string   `json:"executed_amount"`
		RemainingAmount   string   `json:"remaining_amount"`
		ClientOrderID     string   `json:"client_order_id"`
		Options           []string `json:"options"`
		Price             string   `json:"price"`
		OriginalAmount    string   `json:"original_amount"`
	}
)

// NewOrder Creates a new order on Gemini.
// @see https://docs.gemini.com/rest-api/#new-order
func (c *Client) NewOrder(order NewOrderPayload) (resp OrderStatusResponse, err error) {

	order.Nonce = c.createNonce()
	order.Request = "/v1/order/new"

	bodyBytes, err := json.Marshal(order)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/new", bodyBytes, &resp)

	return
}

type (
	CancelOrderPayload struct {
		BaseRequest
		OrderID int `json:"order_id"`
	}
)

// CancelOrder Cancels order given to Gemini.
// @see https://docs.gemini.com/rest-api/#cancel-order
func (c *Client) CancelOrder(order CancelOrderPayload) (resp OrderStatusResponse, err error) {

	order.Nonce = c.createNonce()
	order.Request = "/v1/order/cancel"

	bodyBytes, err := json.Marshal(order)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/cancel", bodyBytes, &resp)

	return
}

type (
	CancelOrdersDetail struct {
		CancelledOrders []int `json:"cancelledOrders"`
		CancelRejects   []int `json:"cancelRejects"`
	}

	CancelOrdersPayload struct {
		Result  string             `json:"result"`
		Details CancelOrdersDetail `json:"details"`
	}
)

// CancelAllSessionOrder Cancels all orders created by session.
// @see https://docs.gemini.com/rest-api/#cancel-all-session-orders
func (c *Client) CancelAllSessionOrder() (resp CancelOrdersPayload, err error) {

	req := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/order/cancel/session",
	}

	bodyBytes, err := json.Marshal(req)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/cancel/session", bodyBytes, &resp)

	return
}

// CancelAllActiveOrders Cancels all orders created by session.
// @see https://docs.gemini.com/rest-api/#cancel-all-active-orders
func (c *Client) CancelAllActiveOrders() (resp CancelOrdersPayload, err error) {

	req := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/order/cancel/all",
	}

	bodyBytes, err := json.Marshal(req)

	if err != nil {
		return
	}

	err = c.Call("POST", "/order/cancel/all", bodyBytes, &resp)

	return
}
