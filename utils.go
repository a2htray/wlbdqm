package wlbdqm

import (
	"errors"
	"log"
	"math"
	"strconv"
)

func DebugPrintln(vs ...interface{}) {
	if appMode == AppModeDebug {
		vs = append([]interface{}{"[debug] "}, vs...)
		log.Println(vs...)
	}
}

func ErrorPrintln(vs ...interface{}) {
	vs = append([]interface{}{"[error] "}, vs...)
	log.Println(vs...)
}

func InfoPrintln(vs ...interface{}) {
	vs = append([]interface{}{"[info] "}, vs...)
	log.Println(vs...)
}

var units = []byte{'K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y'}

func ParseSizeToByte(s string) (float64, error) {
	n := len(s)
	unit := s[n-1]

	var part float64
	for i, v := range units {
		if v == unit {
			part = math.Pow(1024, float64(i+1))
		}
	}

	if part <= 0 {
		return 0, errors.New("unit not found")
	}

	coef, err := strconv.ParseFloat(s[:n-1], 64)
	if err != nil {
		return 0, errors.New("parse size error")
	}

	return coef * part, nil
}
