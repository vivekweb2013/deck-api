package deck

import (
	"github.com/google/uuid"
	"github.com/vivekweb2013/deck-api/internal/apperror"
)

// CreateDeckService represents a service for creating new deck.
//go:generate mockgen -source=create_deck_service.go -package=deck -destination=mock_create_deck_service.go
type CreateDeckService interface {
	CreateDeck(shuffle bool, cards Cards) (*Deck, error)
}

type createDeckService struct {
}

// NewCreateDeckService returns a new service instance of CreateDeckService.
func NewCreateDeckService() CreateDeckService {
	return &createDeckService{}
}

// CreateDeck creates a new deck by accepting the shuffle param & the cards to be included in the deck.
// It returns the newly created deck or any error occurred during deck creation.
func (s *createDeckService) CreateDeck(shuffle bool, cards Cards) (*Deck, error) {
	if !cards.validate() {
		return nil, apperror.NewAppError(apperror.ErrorCodeValidationFailed, "validation failed. invalid cards provided")
	}

	deckID := uuid.NewString()

	deck := Deck{
		ID:       deckID,
		Shuffled: shuffle,
		Cards:    getCards(shuffle, cards),
	}

	decks[deckID] = &deck

	return &deck, nil
}

func getCards(shuffle bool, cards Cards) Cards {
	if len(cards) == 0 {
		cards = fullDeckCards[:]
	}

	if shuffle {
		cards.shuffle()
	}

	return cards
}
