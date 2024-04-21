package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRecordingWinsAndRetrieve(t *testing.T) {
	// Given a server with an in memory player store
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{
		store,
	}

	// When storing 3 wins for a player
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinRequest(player))

	// Then the wins should be 3
	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRetrieveWinsRequest(player))

	assertEqual(t, response.Code, http.StatusOK)
	assertEqual(t, response.Body.String(), "3")
}

func newRecordWinRequest(player string) *http.Request {
	return &http.Request{
		Method: http.MethodPost,
		URL: &url.URL{
			Path: fmt.Sprintf("/player/%s", player),
		},
	}
}

func newRetrieveWinsRequest(player string) *http.Request {
	return &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Path: fmt.Sprintf("/player/%s", player),
		},
	}
}
