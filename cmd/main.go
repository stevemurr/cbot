package main

import (
	"log"
	"os"
	"strconv"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/stevemurr/cbot"
	"github.com/stevemurr/cbot/asset"
	"github.com/stevemurr/cbot/store"
)

func logger(L chan coinbasepro.Message) {
	for {
		select {
		case m := <-L:
			log.Println(m.ProductID, m.Price)
		default:
			continue
		}
	}
}

func main() {
	port, err := strconv.Atoi(os.Getenv("postgres_port"))
	if err != nil {
		panic(err)
	}
	store := store.NewPostgres(
		os.Getenv("postgres_host"),
		os.Getenv("postgres_username"),
		os.Getenv("postgres_password"),
		port,
		"cbot")

	b := cbot.Bot{
		Store: store,
		D:     make(chan coinbasepro.Message),
		L:     make(chan coinbasepro.Message),
		C:     make(chan coinbasepro.Message),
		Assets: map[string]asset.Asset{
			"ADA-USD": &asset.ADA{Name: "ADA-USD"},
		},
		TickerUrl:           "wss://ws-feed.pro.coinbase.com",
		TickerSubscriptions: []string{},
	}

	go b.Listen()
	go b.Evaluate()
	go b.Sinkhole()

	logger(b.L)
}
