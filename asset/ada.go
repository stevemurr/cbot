package asset

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
	// evaluate rules
}
