package httpservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/deck-api/internal/apperror"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

// OpenDeckHandler represents http handler for deck retrival operations.
type OpenDeckHandler struct {
	openDeckService deck.OpenDeckService
}

// NewOpenDeckHandler creates and returns a new handler for serving open-deck endpoint.
func NewOpenDeckHandler(openDeckService deck.OpenDeckService) *OpenDeckHandler {
	return &OpenDeckHandler{openDeckService: openDeckService}
}

// OpenDeck return a given deck by its UUID.
// It returns the deck as a http response.
func (d *OpenDeckHandler) OpenDeck(c *gin.Context) {
	deckID := c.Param("id") // no validation needed as missing id will result in 404 with no call to this handler
	logrus.WithField("deckID", deckID).Info("request to open deck")
	deck, err := d.openDeckService.OpenDeck(deckID)
	if err != nil {
		logrus.WithField("deckID", deckID).Error("request to open deck failed")
		apperror.AbortRequestWithError(c, err)
		return
	}

	deckResp := makeDeckWithCardsResponse(deck)
	c.JSON(http.StatusOK, deckResp)
	logrus.WithField("deckID", deckID).Info("request to open deck successful")
}

func makeDeckWithCardsResponse(d *deck.Deck) *DeckWithCardsResponse {
	deck := makeDeckResponse(d)

	cards := makeCardsResponse(d.Cards)
	r := &DeckWithCardsResponse{
		*deck,
		cards,
	}
	return r
}
