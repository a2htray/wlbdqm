package wlbdqm

import (
	"log"
	"testing"
)

func TestDebugPrintln(t *testing.T) {
	DebugPrintln("qwerty")
	DebugPrintln("123456", "654321")
}

func TestParseSizeToByte(t *testing.T) {
	s := "15.39T"
	v, err := ParseSizeToByte(s)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(v)


	s = "1M"
	v, err = ParseSizeToByte(s)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(v)
}
