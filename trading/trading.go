package trading

import "github.com/preichenberger/go-coinbasepro/v2"

type Tradeable interface {
	Trade(coinbasepro.Message)
}
