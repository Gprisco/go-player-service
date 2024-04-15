package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func TestGETPlayers(t *testing.T) {
	playerServer := &PlayerServer{&StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}}

	t.Run("it should return Pepper's score", func(t *testing.T) {
		// Given
		request := newPlayerRequest("Pepper")
		response := httptest.NewRecorder()

		// When
		playerServer.ServeHTTP(response, request)

		// Then
		assertEqual(t, response.Code, 200)
		assertEqual(t, response.Body.String(), "20")
	})

	t.Run("it should return Floyd's score", func(t *testing.T) {
		// Given
		request := newPlayerRequest("Floyd")
		response := httptest.NewRecorder()

		// When
		playerServer.ServeHTTP(response, request)

		// Then
		assertEqual(t, response.Code, 200)
		assertEqual(t, response.Body.String(), "10")
	})

	t.Run("it should return 404 when the player is not found", func(t *testing.T) {
		// Given
		request := newPlayerRequest("this-player-does-not-exist")
		response := httptest.NewRecorder()

		// When
		playerServer.ServeHTTP(response, request)

		// Then
		assertEqual(t, response.Code, 404)
	})
}

func TestStoreWins(t *testing.T) {
	playerStore := &StubPlayerStore{
		map[string]int{},
	}
	playerServer := &PlayerServer{playerStore}

	t.Run("it should return http 201 when storing a win for a given player", func(t *testing.T) {
		// Given
		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		// When
		playerServer.ServeHTTP(response, request)

		// Then
		assertEqual(t, response.Code, 201)
	})
}

func newPlayerRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertEqual[K comparable](t *testing.T, got K, want K) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
