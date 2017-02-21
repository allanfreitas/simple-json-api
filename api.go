package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

type Stocks []Stock

func handleIndex() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home path")
	})

}

func handleStocks(w http.ResponseWriter, r *http.Request) {
	stocks := Stocks{
		Stock{"PETR4", "Petróleo Brasileiro"},
		Stock{"ELPL3", "Eletropaulo"},
	}
	js, _ := json.Marshal(stocks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func callStocks() http.HandlerFunc {
	return http.HandlerFunc(handleStocks)
}

func main() {

	http.HandleFunc("/", handleIndex())
	http.HandleFunc("/stocks", callStocks())
	http.ListenAndServe(":8080", nil)

}
