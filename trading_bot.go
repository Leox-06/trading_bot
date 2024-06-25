package trading_bot

import (
	"fmt"
	"net/http"
)

type Candle struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
	Time   float64 `json:"time"`
}

type Graph struct {
	Candles  []Candle
	Interval int
	Limit    int
}

type Account struct {
	EUR float64
	BTC float64
}

var client http.Client

func Set_client(c http.Client) {
	client = c
}

var current_price_func func(http.Client) (float64, error)

func Set_current_price_func(f func(http.Client) (float64, error)) {
	current_price_func = f
}

func get_current_price() (float64, error) {
	return current_price_func(client)
}

var graph_func func(http.Client, int, int) (Graph, error)

func Set_graph_func(f func(http.Client, int, int) (Graph, error)) {
	graph_func = f
}

func get_graph(i, l int) (Graph, error) {
	return graph_func(client, i, l)
}

func (account *Account) Buy(volume float64) {
	current_price, err := get_current_price()
	if err != nil {
		fmt.Println(err)
		return
	}
	account.EUR -= volume
	account.BTC += volume / current_price
}

func (account *Account) Sell(volume float64) {
	current_price, err := get_current_price()
	if err != nil {
		fmt.Println(err)
		return
	}

	account.BTC -= volume
	account.EUR += volume * current_price
}
