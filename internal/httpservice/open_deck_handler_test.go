package httpservice

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

func TestOpenDeck(t *testing.T) {
	t.Run("should return a deck when request has valid deck id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockOpenDeckService(ctrl)
		mockService.EXPECT().OpenDeck(deckID).Return(validDeck(false, fullCards[:2]), nil)
		handler := NewOpenDeckHandler(mockService)

		router.GET("/api/v1/decks/:id", handler.OpenDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/decks/%s", deckID), nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `{"deck_id":"6ba3b394-2532-49fd-8739-7b6b89e38d6e", "remaining":2, "shuffled":false, "cards":[{"code":"AC", "suite":"Club", "value":"Ace"}, {"code":"2C", "suite":"Club", "value":"2"}]}`, response.Body.String())
	})

	t.Run("should return error response when retrieving deck fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockOpenDeckService(ctrl)
		mockService.EXPECT().OpenDeck(deckID).Return(nil, errors.New("some error"))
		handler := NewOpenDeckHandler(mockService)

		router.GET("/api/v1/decks/:id", handler.OpenDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/decks/%s", deckID), nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, `{"code":"internal_server_error", "message":"something went wrong."}`, response.Body.String())
	})
}
