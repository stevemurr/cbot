package asset

import (
	"github.com/preichenberger/go-coinbasepro/v2"
)

type Asset interface {
	Listener
	Tradeable
}

type Listener interface {
	// Lookup returns the ticker code
	Lookup() string
}

type Tradeable interface {
	// Evaluate is a goroutine that listens for messages and passes those messages to user logic
	Evaluate(coinbasepro.Message)
}
