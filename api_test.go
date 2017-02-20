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

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, string(body), "Home path")

}
