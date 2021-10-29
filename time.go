package wlbdqm

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	intervalRegex = regexp.MustCompile(`(\d+)([dhms])`)
)

type intervalUnit string

const (
	intervalUnitD intervalUnit = "d"
	intervalUnitH intervalUnit = "h"
	intervalUnitM intervalUnit = "m"
	intervalUnitS intervalUnit = "s"
)

type Interval struct {
	t float64
	unit intervalUnit
}

func (i Interval) String() string {
	return fmt.Sprintf("%v%v", i.t, i.unit)
}

func (i Interval) ToSpec() string {
	return "@every " + i.String()
}

func IntervalFromString(s string) (*Interval, error) {
	if !intervalRegex.MatchString(s) {
		return nil, errors.New("interval format error")
	}

	ss := intervalRegex.FindAllStringSubmatch(s, -1)
	DebugPrintln(ss)

	t, err := strconv.ParseFloat(ss[0][1], 64)
	if err != nil {
		return nil, err
	}

	interval := &Interval{
		t: t,
		unit: intervalUnit(ss[0][2]),
	}

	return interval, nil
}