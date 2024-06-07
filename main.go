package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	indicator "main/indicators"
	"net/http"
)

func main() {

	url := "https://api.youngplatform.com/api/v3/charts?pair=BTC-EUR&interval=60&limit=168"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	var candles []indicator.Candle

	json.Unmarshal(body, &candles)
	graph := indicator.Graph{candles, 60, 168}

	fmt.Println(graph.ATR(14))
}
