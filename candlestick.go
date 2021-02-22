package models

import "time"

// CandleStick represents an OHLC candlestick
type CandleStick struct {
	Time  time.Time `bson:"time"   json:"time,omitempty"`
	Open  float64   `bson:"open"   json:"open,omitempty"`
	High  float64   `bson:"high"   json:"high,omitempty"`
	Low   float64   `bson:"low"    json:"low,omitempty"`
	Close float64   `bson:"close"  json:"close,omitempty"`
}
