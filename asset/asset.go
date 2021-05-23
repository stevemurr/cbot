package asset

import (
	"log"
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
)

type Asset struct {
	Name       string
	actionChan chan *Action
	Tradeable
	Trendable
}

func (a *Asset) Trade(m coinbasepro.Message) {
	someTime := time.Now().Add(time.Second * 10)
	fn := func() {
		log.Println("make some judgments")
		time.Sleep(time.Second * 1)
		log.Println("executed the trade")
	}
	act := NewAction(a, someTime, fn)
	a.actionChan <- act
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
