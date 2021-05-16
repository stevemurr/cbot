package rules

import (
	"github.com/preichenberger/go-coinbasepro/v2"
)

type Rule interface {
	Listener
	Trader
}

type Listener interface {
	// Lookup returns the ticker code
	Lookup() string
}

type Trader interface {
	// Evaluate is a goroutine that listens for messages and passes those messages to user logic
	Evaluate(coinbasepro.Message)
}
