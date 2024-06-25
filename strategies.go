package trading_bot

import "fmt"

func (account *Account) SuperTrend_strategy(length int, factor float64) error{
	graph, err := get_graph(60, 365)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, direction := graph.SuperTrend(length, factor)

	if direction[0] > direction[1] {
		account.Buy(account.EUR)
		fmt.Println("comprato")
	}

	if direction[0] < direction[1] {
		account.Sell(account.BTC)
		fmt.Println("venduto")
	}

	return nil
}
