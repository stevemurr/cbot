package asset

import (
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
)

type Asset struct {
	Name string
	Tradeable
	Trendable
}

func (a *Asset) Trade(m coinbasepro.Message) {

}

func (a *Asset) LastNMinutes(float64) Point {
	return Point{}
}

func (a *Asset) LastLocalMinimum() Point {
	return Point{}
}

func (a *Asset) LastLocalMaximum() Point {
	return Point{}
}

type Tradeable interface {
	Trade(coinbasepro.Message)
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
type Action struct {
	Asset *Asset
	Timer *time.Timer
}

func NewAction(asset *Asset, when time.Time, fn func()) *Action {
	return &Action{
		Asset: asset,
		Timer: time.AfterFunc(time.Since(when), fn),
	}
}
