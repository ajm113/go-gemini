package gemini

import "encoding/json"

type (
	NationalVolumeDayVolume struct {
		Date           string  `json:"date"`
		NationalVolume float64 `json:"notional_volume"`
	}

	NationalVolumeResponse struct {
		Date          string `json:"date"`
		LastUpdatedMS int64  `json:"last_updated_ms"`

		WebMakerFeeBPS  int `json:"web_maker_fee_bps"`
		WebTakerFeeBPS  int `json:"web_taker_fee_bps"`
		WebActionFeeBPS int `json:"web_auction_fee_bps"`

		APIMakerFeeBPS  int `json:"api_maker_fee_bps"`
		APITakerFeeBPS  int `json:"api_taker_fee_bps"`
		APIActionFeeBPS int `json:"api_auction_fee_bps"`

		FixMakerFeeBPS  int `json:"fix_maker_fee_bps"`
		FixTakerFeeBPS  int `json:"fix_taker_fee_bps"`
		FixActionFeeBPS int `json:"fix_auction_fee_bps"`

		BlockMakerFeeBPS int `json:"block_maker_fee_bps"`
		BlockTakerFeeBPS int `json:"block_taker_fee_bps"`

		National30DVolume float64                   `json:"notional_30d_volume"`
		National1DVolume  []NationalVolumeDayVolume `json:"notional_1d_volume"`
	}
)

// GetNationalVolume The response will be a JSON object representing the volume in price currency that has been traded across all pairs over a period of 30 days.
// @see https://docs.gemini.com/rest-api/#get-notional-volume
func (c *Client) GetNationalVolume() (resp NationalVolumeResponse, err error) {

	body := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/notionalvolume",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/notionalvolume", bodyBytes, &resp)
	return
}

type (
	TradeVolume struct {
		Symbol           string `json:"symbol"`
		BaseCurrency     string `json:"base_currency"`
		NationalCurrency string `json:"notional_currency"`

		DataDate          string  `json:"data_date"`
		TotalVolumeBase   float64 `json:"total_volume_base"`
		MakerBuySellRatio float64 `json:"maker_buy_sell_ratio"`

		BuyMakerBase     float64 `json:"buy_maker_base"`
		BuyMakerNotional float64 `json:"buy_maker_notional"`
		BuyMakerCount    float64 `json:"buy_maker_count"`

		SellMakerBase     float64 `json:"sell_maker_base"`
		SellMakerNotional float64 `json:"sell_maker_notional"`
		SellMakerCount    float64 `json:"sell_maker_count"`

		BuyTakerBase     float64 `json:"buy_taker_base"`
		BuyTakerNotional float64 `json:"buy_taker_notional"`
		BuyTakerCount    float64 `json:"buy_taker_count"`

		SellTakerBase     float64 `json:"sell_taker_base"`
		SellTakerNotional float64 `json:"sell_taker_notional"`
		SellTakerCount    float64 `json:"sell_taker_count"`
	}

	TradeVolumeResponse [][]TradeVolume
)

// GetNationalVolume The response will be a JSON object representing the volume in price currency that has been traded across all pairs over a period of 30 days.
// @see https://docs.gemini.com/rest-api/#get-trade-volume
func (c *Client) GetTradeVolume() (resp TradeVolumeResponse, err error) {

	body := BaseRequest{
		Nonce:   c.createNonce(),
		Request: "/v1/tradevolume",
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	err = c.Call("POST", "/tradevolume", bodyBytes, &resp)
	return
}
