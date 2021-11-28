package asset

import (
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
)

type Asset struct {
	Name string `json:"name"`
	Trendable
}

func (a *Asset) Trade(m coinbasepro.Message) {}

func (a *Asset) LastNMinutes(float64) Point {
	return Point{}
}

func (a *Asset) LastLocalMinimum() Point {
	return Point{}
}

func (a *Asset) LastLocalMaximum() Point {
	return Point{}
}

type Trendable interface {
	LastNMinutes(float64) Point
	LastLocalMinimum() Point
	LastLocalMaximum() Point
}

type Point struct {
	Time     time.Time
	Average  float64
	Minimum  float64
	Maximum  float64
	STD      float64
	Variance float64
}
