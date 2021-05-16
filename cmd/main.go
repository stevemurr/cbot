package main

import (
	"log"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/stevemurr/cbot"
	"github.com/stevemurr/cbot/rules"
	"github.com/stevemurr/cbot/store"
)

func main() {
	store := store.NewPostgres("192.168.1.99", "murr", "flag9012", 5432, "cbot")

	b := cbot.Bot{
		Store: store,
		D:     make(chan coinbasepro.Message),
		L:     make(chan coinbasepro.Message),
		C:     make(chan coinbasepro.Message),
		Rules: map[string]rules.Rule{
			"ADA-USD": &rules.ADA{Name: "ADA-USD"},
		},
		TickerUrl:           "wss://ws-feed.pro.coinbase.com",
		TickerSubscriptions: []string{},
	}

	go b.Listen()
	go b.Evaluate()
	go b.Sinkhole()
	for {
		select {
		case m := <-b.L:
			log.Println(m.ProductID, m.Price)
		default:
			continue
		}
	}

}
