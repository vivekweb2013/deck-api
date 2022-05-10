package httpservice

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/deck-api/internal/apperror"
	"github.com/vivekweb2013/deck-api/internal/deck"
)

// DrawCardsHandler represents http handler for card-draw operations.
type DrawCardsHandler struct {
	drawCardsService deck.DrawCardsService
}

// NewDrawCardsHandler creates and returns a new handler for serving card-draw endpoint.
func NewDrawCardsHandler(drawCardsService deck.DrawCardsService) *DrawCardsHandler {
	return &DrawCardsHandler{drawCardsService: drawCardsService}
}

// DrawCards draws card(s) from the requested deck.
// It returns the drawn cards.
func (d *DrawCardsHandler) DrawCards(c *gin.Context) {
	deckID := c.Param("id")
	if err := validation.Validate(deckID, validation.Required); err != nil {
		apperror.AbortRequestWithError(c, apperror.NewAppError(apperror.ErrorCodeValidationFailed, fmt.Sprintf("deckID: %s", err.Error())))
		return
	}
	countParam := c.Query("count")
	if err := validation.Validate(countParam, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{1,2}$"))); err != nil {
		apperror.AbortRequestWithError(c, apperror.NewAppError(apperror.ErrorCodeValidationFailed, fmt.Sprintf("count: %s", err.Error())))
		return
	}
	count, _ := strconv.Atoi(countParam)
	logrus.WithField("deckID", deckID).WithField("count", count).Info("request to draw card(s)")
	cards, err := d.drawCardsService.DrawCards(deckID, count)
	if err != nil {
		logrus.WithField("deckID", deckID).WithField("count", count).Error("request to draw card(s) failed")
		apperror.AbortRequestWithError(c, err)
		return
	}

	cardsResp := makeCardsResponse(cards)
	c.JSON(http.StatusOK, cardsResp)
	logrus.WithField("deckID", deckID).WithField("count", count).Info("request to draw card(s) successful")
}

func makeCardsResponse(cards deck.Cards) []CardResponse {
	value := map[string]string{"A": "Ace", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9", "10": "10", "J": "Jack", "Q": "Queen", "K": "King"}
	suite := map[string]string{"H": "Heart", "D": "Diamond", "C": "Club", "S": "Spade"}
	cardsResponse := []CardResponse{}
	for _, c := range cards {
		cardsResponse = append(cardsResponse, CardResponse{
			Value: value[string(c[:len(c)-1])],
			Suite: suite[string(c[len(c)-1:])],
			Code:  c,
		})
	}
	return cardsResponse
}
