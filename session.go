package gemini

import "encoding/json"

type (
	HeartBeatResponse struct {
		Result string `json:"result"`
	}
)

// HeartBeat This will prevent a session from timing out and canceling orders if the require heartbeat flag has been set. Note that this is only required if no other private API requests have been made. The arrival of any message resets the heartbeat timer.
// @see https://docs.gemini.com/rest-api/#heartbeat
func (c *Client) HeartBeat() (resp HeartBeatResponse, err error) {

	body := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/heartbeat",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/heartbeat", bodyBytes, &resp)
	return
}
