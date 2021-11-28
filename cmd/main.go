package main

import (
	"log"
	"os"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/stevemurr/cbot"
	"github.com/stevemurr/cbot/config"
	"github.com/stevemurr/cbot/store"
)

func logger(L chan coinbasepro.Message) {
	for {
		select {
		case m := <-L:
			log.Println(m.ProductID, m.Side, m.Price)
		default:
			continue
		}
	}
}

func main() {

	store := store.NewPostgres(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	c, err := config.ReadConfig("./config.json")
	if err != nil {
		panic(err)
	}

	b := cbot.Bot{
		Store:               store,
		D:                   make(chan coinbasepro.Message),
		L:                   make(chan coinbasepro.Message),
		C:                   make(chan coinbasepro.Message),
		Assets:              c.AssetMap,
		TickerUrl:           "wss://ws-feed.exchange.coinbase.com",
		TickerSubscriptions: []string{},
	}

	go b.ListenToTicker()
	go b.Evaluate()
	go b.LogToDatabase()

	logger(b.L)
}
