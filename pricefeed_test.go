package gemini

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPriceFeed(t *testing.T) {
	priceFeed, err := testClient.GetPriceFeed()
	assert.Nil(t, err)
	assert.NotNil(t, priceFeed)
	assert.True(t, len(priceFeed) > 0)

	for _, c := range priceFeed {
		assert.True(t, c.Pair != "")
		assert.True(t, c.PercentChange24h != "")
		assert.True(t, c.Price != "")
	}
}
