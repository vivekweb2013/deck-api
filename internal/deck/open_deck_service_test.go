package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenDeck(t *testing.T) {
	t.Run("should return error when deck with given id does not exist", func(t *testing.T) {
		service := NewOpenDeckService()
		deckID := "6ba3b394-2532-49fd-8739-7b6b89e38d6e"

		_, err := service.OpenDeck(deckID)

		assert.EqualError(t, err, "error code: invalid_request, message: can not open deck. the deck with given id does not exist")
	})
}
