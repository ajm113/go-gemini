package gemini

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNationalVolume(t *testing.T) {
	nationalVolumeResponse, err := testClient.GetNationalVolume()
	assert.Nil(t, err)
	assert.NotNil(t, nationalVolumeResponse)
}

func TestTradeVolume(t *testing.T) {
	tradeVolumeResponse, err := testClient.GetTradeVolume()
	assert.Nil(t, err)
	assert.NotNil(t, tradeVolumeResponse)
	assert.True(t, len(tradeVolumeResponse) > 0)
}
