package main

import (
	"encoding/json"
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
	port, err := strconv.Atoi(os.Getenv("cbot_postgres_port"))
	if err != nil {
		panic(err)
	}
	store := store.NewPostgres(
		os.Getenv("cbot_postgres_host"),
		os.Getenv("cbot_postgres_username"),
		os.Getenv("cbot_postgres_password"),
		port,
		"cbot")

	assets := map[string]asset.Asset{}
	json.Unmarshal([]byte(string(os.Getenv("cbot_assets"))), &assets)

	b := cbot.Bot{
		Store:               store,
		D:                   make(chan coinbasepro.Message),
		L:                   make(chan coinbasepro.Message),
		C:                   make(chan coinbasepro.Message),
		Assets:              assets,
		TickerUrl:           "wss://ws-feed.pro.coinbase.com",
		TickerSubscriptions: []string{},
	}

	go b.Listen()
	go b.Evaluate()
	go b.Sinkhole()

	logger(b.L)
}
