package deck

import (
	"github.com/vivekweb2013/deck-api/internal/apperror"
)

// DrawCardsService represents a service to draw cards from a deck.
//go:generate mockgen -source=draw_cards_service.go -package=deck -destination=mock_draw_cards_service.go
type DrawCardsService interface {
	DrawCards(deckID string, count int) (Cards, error)
}

type drawCardsService struct {
}

// NewDrawCardsService returns a new service instance of DrawCardsService.
func NewDrawCardsService() DrawCardsService {
	return &drawCardsService{}
}

// DrawCards draws cards from a deck of given deckID. The count indicates number of cards to draw.
// It returns the drawn cards or any error occurred while drawing cards.
func (s *drawCardsService) DrawCards(deckID string, count int) (Cards, error) {
	d, exist := decks[deckID]
	if !exist {
		return nil, apperror.NewAppError(apperror.ErrorCodeInvalidRequest, "can not draw cards from deck. the deck with given id does not exist")
	}

	if count+d.DrawCount > len(d.Cards) {
		return nil, apperror.NewAppError(apperror.ErrorCodeInvalidRequest, "can not draw cards. not enough cards present inside deck")
	}

	d.DrawCount += count

	return d.Cards[d.DrawCount-count : d.DrawCount], nil
}
