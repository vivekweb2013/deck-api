package httpservice

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/deck-api/internal/apperror"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

// CreateDeckHandler represents http handler for deck creation operations.
type CreateDeckHandler struct {
	createDeckService deck.CreateDeckService
}

// NewCreateDeckHandler creates and returns a new handler for serving create deck endpoint.
func NewCreateDeckHandler(createDeckService deck.CreateDeckService) *CreateDeckHandler {
	return &CreateDeckHandler{createDeckService: createDeckService}
}

// CreateDeck create a new deck. It accepts shuffle and cards query params.
// It returns the deck as a http response.
func (d *CreateDeckHandler) CreateDeck(c *gin.Context) {
	shuffle := c.Query("shuffle") == "true"
	cardsParam := c.Query("cards")

	cards := []string{}
	if cardsParam != "" {
		cards = strings.Split(cardsParam, ",")
	}

	logrus.WithField("shuffle", shuffle).WithField("cards", cards).Info("request to create a new deck")

	deck, err := d.createDeckService.CreateDeck(shuffle, cards)
	if err != nil {
		logrus.WithField("shuffle", shuffle).WithField("cards", cards).Error("request to create a new deck failed")
		apperror.AbortRequestWithError(c, err)
		return
	}
	deckResp := makeDeckResponse(deck)
	c.JSON(http.StatusOK, deckResp)
	logrus.WithField("shuffle", shuffle).WithField("cards", cards).Info("request to create a new deck successful")
}

func makeDeckResponse(d *deck.Deck) *DeckResponse {
	return &DeckResponse{
		DeckID:    d.ID,
		Shuffled:  d.Shuffled,
		Remaining: len(d.Cards),
	}
}
