package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	handleIndex := handleIndex()

	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	handleIndex.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Home path", string(body))

}

func TestStocks(t *testing.T) {

	handler := callStocks()

	req, _ := http.NewRequest("GET", "/stocks", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, "[{\"ticker\":\"PETR4\",\"name\":\"Petróleo Brasileiro\"},{\"ticker\":\"ELPL3\",\"name\":\"Eletropaulo\"}]", string(body))
}
