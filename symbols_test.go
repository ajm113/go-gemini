package gemini

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSymbols(t *testing.T) {
	symbolsResponse, err := testClient.GetSymbols()
	assert.Nil(t, err)
	assert.NotNil(t, symbolsResponse)
	assert.True(t, len(symbolsResponse) > 1)
}

func TestSymbolDetails(t *testing.T) {
	symbolDetails, err := testClient.GetSymbolDetails("btcusd")
	assert.Nil(t, err)
	assert.NotNil(t, symbolDetails)
	assert.True(t, symbolDetails.TickSize == 1e-8)
	assert.True(t, symbolDetails.Status == "open")
	assert.True(t, symbolDetails.BaseCurrency == "BTC")
	assert.True(t, symbolDetails.QuoteCurrency == "USD")
	assert.True(t, symbolDetails.Symbol == "BTCUSD")
}
