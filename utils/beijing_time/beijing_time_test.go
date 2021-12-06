package beijing_time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	now := Now()
	expectedOffset := 8 * 60 * 60
	_, offset := now.Zone()

	assert.Equal(t, expectedOffset, offset, "time zone error")
}

func TestStartTimeOfDay(t *testing.T) {
	time, err := time.Parse(time.RFC3339, "2021-07-08T01:02:03+08:00")
	startTimeOfDay := StartTimeOfDay(time)

	assert := assert.New(t)
	assert.Nil(err, "time format error")
	assert.Equal(2021, startTimeOfDay.Year(), "year should be equal")
	assert.EqualValues(7, startTimeOfDay.Month(), "month should be equal")
	assert.Equal(8, startTimeOfDay.Day(), "day should be equal")
	assert.Zero(startTimeOfDay.Hour(), "hour should be zero")
	assert.Zero(startTimeOfDay.Minute(), "minute should be zero")
	assert.Zero(startTimeOfDay.Second(), "second should be zero")
	assert.Zero(startTimeOfDay.Nanosecond(), "nanosecond should be zero")
}
