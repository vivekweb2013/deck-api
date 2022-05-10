package httpservice

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

var fullCards = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
}

const deckID = "6ba3b394-2532-49fd-8739-7b6b89e38d6e"

func TestCreateDeck(t *testing.T) {
	t.Run("should return a new full-deck when shuffle and cards params missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := getRouter()
		mockService := deck.NewMockCreateDeckService(ctrl)
		shuffle := false
		cards := deck.Cards{}
		mockService.EXPECT().CreateDeck(shuffle, cards).Return(getDeck(shuffle, fullCards[:]), nil)
		handler := NewCreateDeckHandler(mockService)

		router.POST("/api/v1/decks", handler.CreateDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `{"deck_id":"6ba3b394-2532-49fd-8739-7b6b89e38d6e", "remaining":52, "shuffled":false}`, response.Body.String())
	})

	t.Run("should return a new shuffled full-deck when shuffle param is passed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := getRouter()
		mockService := deck.NewMockCreateDeckService(ctrl)
		shuffle := true
		cards := deck.Cards{}
		mockService.EXPECT().CreateDeck(shuffle, cards).Return(getDeck(shuffle, fullCards[:]), nil)
		handler := NewCreateDeckHandler(mockService)

		router.POST("/api/v1/decks", handler.CreateDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks?shuffle=true", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `{"deck_id":"6ba3b394-2532-49fd-8739-7b6b89e38d6e", "remaining":52, "shuffled":true}`, response.Body.String())
	})

	t.Run("should return a new partial deck when cards param is passed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := getRouter()
		mockService := deck.NewMockCreateDeckService(ctrl)
		shuffle := false
		cards := deck.Cards{"5H", "JS"}
		mockService.EXPECT().CreateDeck(shuffle, cards).Return(getDeck(shuffle, cards), nil)
		handler := NewCreateDeckHandler(mockService)

		router.POST("/api/v1/decks", handler.CreateDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks?cards=5H,JS", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `{"deck_id":"6ba3b394-2532-49fd-8739-7b6b89e38d6e", "remaining":2, "shuffled":false}`, response.Body.String())
	})

	t.Run("should return a new partial & shuffled deck when shuffle & cards param is passed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := getRouter()
		mockService := deck.NewMockCreateDeckService(ctrl)
		shuffle := true
		cards := deck.Cards{"2H", "9D"}
		mockService.EXPECT().CreateDeck(shuffle, cards).Return(getDeck(shuffle, cards), nil)
		handler := NewCreateDeckHandler(mockService)

		router.POST("/api/v1/decks", handler.CreateDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks?shuffle=true&cards=2H,9D", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.JSONEq(t, `{"deck_id":"6ba3b394-2532-49fd-8739-7b6b89e38d6e", "remaining":2, "shuffled":true}`, response.Body.String())
	})

	t.Run("should return error response when creating new deck fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := getRouter()
		mockService := deck.NewMockCreateDeckService(ctrl)
		mockService.EXPECT().CreateDeck(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
		handler := NewCreateDeckHandler(mockService)

		router.POST("/api/v1/decks", handler.CreateDeck)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/decks", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.JSONEq(t, `{"code":"internal_server_error", "message":"something went wrong."}`, response.Body.String())
	})
}

func getRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.UseRawPath = true
	return router
}

func getDeck(shuffle bool, cards deck.Cards) *deck.Deck {
	deck := &deck.Deck{
		ID:       deckID,
		Shuffled: shuffle,
		Cards:    cards,
	}
	return deck
}
