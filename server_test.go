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
		assertEqual(t, response.Body.String(), "20")
	})

	t.Run("it should return Floyd's score", func(t *testing.T) {
		// Given
		request := newPlayerRequest("Floyd")
		response := httptest.NewRecorder()

		// When
		playerServer.ServeHTTP(response, request)

		// Then
		assertEqual(t, response.Body.String(), "10")
	})
}

func newPlayerRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertEqual(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
