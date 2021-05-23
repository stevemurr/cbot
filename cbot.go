package cbot

import (
	"log"

	ws "github.com/gorilla/websocket"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
	"github.com/stevemurr/cbot/asset"
	"github.com/stevemurr/cbot/store"
)

type Bot struct {
	wsConn              *ws.Conn
	TickerUrl           string
	TickerSubscriptions []string
	D                   chan coinbasepro.Message
	C                   chan coinbasepro.Message
	L                   chan coinbasepro.Message
	Assets              map[string]asset.Asset
	Store               store.Store
}

func (b *Bot) getTickerConnection() error {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(b.TickerUrl, nil)
	if err != nil {
		return err
	}

	subscribe := coinbasepro.Message{
		Type: "subscribe",
		Channels: []coinbasepro.MessageChannel{
			{
				Name:       "ticker",
				ProductIds: b.TickerSubscriptions,
			},
		},
	}
	if err := wsConn.WriteJSON(subscribe); err != nil {
		return err
	}
	b.wsConn = wsConn
	return nil
}

func (b *Bot) updateSubscriptions() {
	subs := []string{}
	for sub := range b.Assets {
		subs = append(subs, sub)
	}
	b.TickerSubscriptions = subs
}

// ListenToTicker --
func (b *Bot) Listen() error {
	b.updateSubscriptions()

	if err := b.getTickerConnection(); err != nil {
		return err
	}
	for {
		message := coinbasepro.Message{}
		if err := b.wsConn.ReadJSON(&message); err != nil {
			return err
		}
		b.C <- message
		b.L <- message
		b.D <- message
	}
}

func (b *Bot) Evaluate() error {
	for {
		select {
		case m := <-b.C:
			rule := b.Assets[m.ProductID]
			rule.Trade(m)
		default:
			continue
		}
	}
}

func (b *Bot) Sinkhole() error {
	if err := b.Store.Connect(); err != nil {
		return err
	}
	for {
		select {
		case m := <-b.D:
			if err := b.Store.Insert(m.Price, "now()", m.ProductID); err != nil {
				log.Println(err)
			}
		default:
			continue
		}
	}
}
