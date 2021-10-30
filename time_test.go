package wlbdqm

import (
	"testing"
)

func TestIntervalFromString(t *testing.T) {
	interval, err := IntervalFromString("1h")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(interval)

	interval, err = IntervalFromString("12d")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(interval)

	interval, err = IntervalFromString("12m")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(interval)

}
