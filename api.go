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

type Stocks []Stock

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

func main() {

	router := httprouter.New()

	router.GET("/", handleIndex)
	router.GET("/stocks", handleStocks)

	http.ListenAndServe(":8080", nil)

}
