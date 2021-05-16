package rules

import (
	"github.com/preichenberger/go-coinbasepro/v2"
)

type ADA struct {
	Name string
}

func (a *ADA) Lookup() string {
	return a.Name
}

func (a *ADA) Evaluate(m coinbasepro.Message) {
	// if price goes 10% above last known local maximum
	//   do shit
	// else
	//   hold
	// if price goes 10% last known local minimum
	//   look for gainer and move
}
