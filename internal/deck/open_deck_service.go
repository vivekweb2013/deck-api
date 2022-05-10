package deck

import (
	"sync"

	"github.com/vivekweb2013/deck-api/internal/apperror"
)

// OpenDeckService represents a service to open a deck.
//go:generate mockgen -source=open_deck_service.go -package=deck -destination=mock_open_deck_service.go
type OpenDeckService interface {
	OpenDeck(deckID string) (*Deck, error)
}

type openDeckService struct {
}

// NewOpenDeckService returns a new service instance of OpenDeckService.
func NewOpenDeckService() OpenDeckService {
	return &openDeckService{}
}

// OpenDeck retrieves a deck by it's id.
// It returns the deck or any error occurred while retrieving it.
func (s *openDeckService) OpenDeck(deckID string) (*Deck, error) {
	d, exist := decks[deckID]
	if !exist {
		return nil, apperror.NewAppError(apperror.ErrorCodeInvalidRequest, "can not open deck. the deck with given id does not exist")
	}
	return &Deck{
		ID:        d.ID,
		Shuffled:  d.Shuffled,
		Cards:     d.Cards[d.DrawCount:],
		DrawCount: d.DrawCount,
		Mutex:     sync.Mutex{},
	}, nil
}
