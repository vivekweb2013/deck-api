package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fullCards = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
}

func TestCreateDeck(t *testing.T) {
	t.Run("should return a new ordered full-deck when shuffle is false and cards empty", func(t *testing.T) {
		service := NewCreateDeckService()
		shuffle := false
		cards := Cards{}

		d, err := service.CreateDeck(shuffle, cards)

		assert.NoError(t, err)
		assert.NotEmpty(t, d.ID)
		assert.Equal(t, shuffle, d.Shuffled)
		assert.EqualValues(t, fullCards[:], d.Cards)
	})

	t.Run("should return a new ordered partial deck when shuffle is false and cards provided", func(t *testing.T) {
		service := NewCreateDeckService()
		shuffle := false
		cards := Cards{"3C", "7H"}

		d, err := service.CreateDeck(shuffle, cards)

		assert.NoError(t, err)
		assert.NotEmpty(t, d.ID)
		assert.Equal(t, shuffle, d.Shuffled)
		assert.EqualValues(t, cards, d.Cards)
	})

	t.Run("should return a new shuffled full-deck when shuffle is true and cards empty", func(t *testing.T) {
		service := NewCreateDeckService()
		shuffle := true
		cards := Cards{}

		d, err := service.CreateDeck(shuffle, cards)

		assert.NoError(t, err)
		assert.NotEmpty(t, d.ID)
		assert.Equal(t, shuffle, d.Shuffled)
		// Verify that all the cards are present in deck
		assert.ElementsMatch(t, fullCards[:], d.Cards)
		// Verify that the order is not the same
		assert.NotEqualValues(t, fullCards[:], d.Cards)
	})

	t.Run("should return error when invalid cards provided", func(t *testing.T) {
		service := NewCreateDeckService()
		shuffle := false
		cards := Cards{"3Z", "7Y"}

		_, err := service.CreateDeck(shuffle, cards)

		assert.EqualError(t, err, "validation failed. invalid cards provided")
	})
}
