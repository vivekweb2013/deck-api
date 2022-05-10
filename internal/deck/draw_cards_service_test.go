package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawCards(t *testing.T) {
	t.Run("should return error when deck with given id does not exist", func(t *testing.T) {
		service := NewDrawCardsService()
		deckID := "6ba3b394-2532-49fd-8739-7b6b89e38d6e"
		count := 2

		_, err := service.DrawCards(deckID, count)

		assert.EqualError(t, err, "error code: invalid_request, message: can not draw cards from deck. the deck with given id does not exist")
	})
}
