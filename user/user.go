// package user represents a single users actions and state
package user

type User interface {
	Trader
	Ledger
	Entity
}

type Entity interface {
	GetName() string
}

type Ledger interface {
	GetLedger() map[string]string
}

type Trader interface {
	Trade()
}
