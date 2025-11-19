package seqx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {
	id := NewID()
	ti, err := GetTime(id)
	assert.NoError(t, err)

	// delta no more than 1s
	assert.Equal(t, 0, int(time.Until(ti).Seconds()))
}
