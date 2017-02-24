package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

type Data struct {
	Date  string `json:"data"`
	Value string `json:"value"`
}

type Stocks []Stock

type StockData []Data

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Home path")

}

func handleStocks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stocks := Stocks{
		Stock{"PETR4", "Petróleo Brasileiro"},
		Stock{"ELPL3", "Eletropaulo"},
	}
	js, _ := json.Marshal(stocks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func handleStockData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	stockData := StockData{
		Data{"01/01/2017", "39.40"},
		Data{"02/01/2017", "39.00"},
	}
	tempMap := make(map[string]string)
	for _, stock := range stockData {
		tempMap[stock.Date] = stock.Value
	}
	js, _ := json.Marshal(tempMap)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {

	router := httprouter.New()

	router.GET("/", handleIndex)
	router.GET("/stocks", handleStocks)
	router.GET("/stocks/:ticker", handleStockData)

	http.ListenAndServe(":8080", nil)

}
