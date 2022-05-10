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

func TestDrawCards(t *testing.T) {
	t.Run("should return drawn cards when request has deck-id and count params", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockDrawCardsService(ctrl)
		count := 2
		cards := fullCards[:2]
		mockService.EXPECT().DrawCards(deckID, count).Return(cards, nil)
		handler := NewDrawCardsHandler(mockService)

		router.POST("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/decks/%s/draw?count=2", deckID), nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `[{"code":"AC", "suite":"Club", "value":"Ace"}, {"code":"2C", "suite":"Club", "value":"2"}]`, response.Body.String())
	})

	t.Run("should return error response when card draw operation fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockDrawCardsService(ctrl)
		mockService.EXPECT().DrawCards(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
		handler := NewDrawCardsHandler(mockService)

		router.POST("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/decks/%s/draw?count=2", deckID), nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, `{"code":"internal_server_error", "message":"something went wrong."}`, response.Body.String())
	})

	t.Run("should return validation error response when request does not have deck id param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockDrawCardsService(ctrl)
		handler := NewDrawCardsHandler(mockService)

		router.POST("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks//draw?count=2", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, `{"code":"validation_failed", "message":"deckID: cannot be blank"}`, response.Body.String())
	})

	t.Run("should return validation error response when request does not have count param", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := router()
		mockService := deck.NewMockDrawCardsService(ctrl)
		handler := NewDrawCardsHandler(mockService)

		router.POST("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/decks/%s/draw", deckID), nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.JSONEq(t, `{"code":"validation_failed", "message":"count: cannot be blank"}`, response.Body.String())
	})
}
