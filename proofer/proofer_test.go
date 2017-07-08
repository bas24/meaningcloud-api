package proofer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	apiKeyCorrect := "123ewq456Kre8910"
	must := apiKeyCorrect
	key, _ := New(apiKeyCorrect)
	assert.Equal(t, key.MeaningcloudKey, must, "Must return correct api key")
}
