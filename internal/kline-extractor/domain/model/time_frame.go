package model

import (
	"fmt"
	"strconv"
)

// TimeFrame represents time frame of a chart
type TimeFrame string

const (
	ONE_MIN        = TimeFrame("1")
	THREE_MIN      = TimeFrame("3")
	FIVE_MIN       = TimeFrame("5")
	FIFTEEN_MIN    = TimeFrame("15")
	THIRTY_MIN     = TimeFrame("30")
	FORTY_FIVE_MIN = TimeFrame("45")
	ONE_HOUR       = TimeFrame("60")
	TWO_HOUR       = TimeFrame("120")
	THREE_HOUR     = TimeFrame("180")
	FOUR_HOUR      = TimeFrame("240")
	ONE_DAY        = TimeFrame("D")
	ONE_WEEK       = TimeFrame("W")
	ONE_MONTH      = TimeFrame("M")
)

// ToSecond converts given time frame to equal seconds of that time frame's period
func (t TimeFrame) ToSecond() (int64, error) {
	switch t {
	case ONE_DAY:
		return 24 * 60 * 60, nil
	case ONE_WEEK:
		return 7 * 24 * 60 * 60, nil
	case ONE_MONTH:
		return 30 * 24 * 60 * 60, nil
	}
	val, err := strconv.ParseInt(string(t), 10, 64)
	if err != nil || val == 0 {
		return 0, fmt.Errorf("given timeframe is not correct")
	}
	return val * 60, nil
}
