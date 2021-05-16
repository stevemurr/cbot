package rules

import "github.com/preichenberger/go-coinbasepro/v2"

type BTC struct {
	Name string
}

func (b BTC) Lookup() string {
	return b.Name
}

func (b BTC) Evaluate(c chan coinbasepro.Message) {

}
