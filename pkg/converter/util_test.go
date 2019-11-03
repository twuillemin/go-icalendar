package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadPositiveHourMinute(t *testing.T) {
	hourMinute, err := readHourMinute("+1234")
	assert.NoError(t, err)
	assert.Equal(t, 12, hourMinute.Hour)
	assert.Equal(t, 34, hourMinute.Minute)
}

func TestReadUnsignedHourMinute(t *testing.T) {
	hourMinute, err := readHourMinute("1234")
	assert.NoError(t, err)
	assert.Equal(t, 12, hourMinute.Hour)
	assert.Equal(t, 34, hourMinute.Minute)
}

func TestReadNegativeHourMinute(t *testing.T) {
	hourMinute, err := readHourMinute("-1234")
	assert.NoError(t, err)
	assert.Equal(t, -12, hourMinute.Hour)
	assert.Equal(t, 34, hourMinute.Minute)
}
