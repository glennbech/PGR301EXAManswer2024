package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Record struct {
	Side     string  `json:"side"`
	Price    float64 `json:"price,string"`
	ID       string  `json:"id"`
	Quantity float64 `json:"quantity,string"`
}

type ConversionRate struct {
	Rates map[string]float64 `json:"rates"`
}

/*
    Koden viser den høyeste prisen noen er villig til å betale, og den laveste prisen noen er villige for å selge
    BitCoin for å kryptobørsen NBX. Koden ser bare å første "page" i responsen. Differansnen kalles gjerne "Spread"
    dokumentasjon https://app.nbx.com/developers#tag/Order-Book
*/

func main() {
	marketID := "BTC-NOK" // Replace YOUR_MARKET_ID with the actual market ID
	ordersURL := fmt.Sprintf("https://api.nbx.com/markets/%s/orders", marketID)

	resp, err := http.Get(ordersURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var records []Record
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	if err != nil {
		fmt.Println("Error fetching orders:", err)
		return
	}
	var highestBuy, lowestSell float64 = 0, math.MaxFloat64
	for _, record := range records {
		if record.Side == "BUY" && record.Price > highestBuy {
			highestBuy = record.Price
		} else if record.Side == "SELL" && record.Price < lowestSell {
			lowestSell = record.Price
		}
	}
	fmt.Printf("Highest BUY: %.2f\nLowest SELL: %.2f\n", highestBuy/10.42, lowestSell/10.42)
}
