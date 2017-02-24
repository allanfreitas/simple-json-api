package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {
		handleIndex(w, r, nil)
	}

	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Home path", string(body))

}

func TestStocks(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		handleStocks(w, r, nil)
	}

	req, _ := http.NewRequest("GET", "/stocks", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "[{\"ticker\":\"PETR4\",\"name\":\"Petróleo Brasileiro\"},{\"ticker\":\"ELPL3\",\"name\":\"Eletropaulo\"}]", string(body))
}

func TestStocksData(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request, param string) {
		p := httprouter.Params{
			httprouter.Param{"ticker", param},
		}
		handleStockData(w, r, p)
	}

	req, _ := http.NewRequest("GET", "/stocks/PETR4", nil)
	w := httptest.NewRecorder()

	handler(w, req, "PETR4")

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "{\"01/01/2017\":\"39.40\",\"02/01/2017\":\"39.00\"}", string(body))

}
