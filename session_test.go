package gemini

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestHearBeart(t *testing.T) {
	heartBeatResponse, err := testClient.HeartBeat()
	assert.Nil(t, err)

	assert.True(t, heartBeatResponse.Result == "ok")
}
