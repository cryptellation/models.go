package models

import (
	"testing"
	"time"
)

func TestCandleStickEqual(t *testing.T) {
	inst, err := time.Parse("02/01/2006", "10/06/2021")
	if err != nil {
		t.Fatal(err)
	}

	a := CandleStick{inst, 0, 1, 2, 3}
	b := CandleStick{inst, 0, 1, 2, 3}

	if a.Equal(&b) == false {
		t.Error("Candlesticks should be equal")
	}
}

func TestCandleStickNotEqual(t *testing.T) {
	inst, err := time.Parse("02/01/2006", "10/06/2021")
	if err != nil {
		t.Fatal(err)
	}

	a := CandleStick{inst, 0, 1, 2, 3}
	b := CandleStick{inst, 2, 4, 2, 3}

	if a.Equal(&b) {
		t.Error("Candlesticks should not be equal")
	}
}
