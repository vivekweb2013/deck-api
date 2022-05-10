package deck

import (
	"math/rand"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

var fullDeckCards = [52]string{
	"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
	"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
	"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
	"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
}

type Cards []string

type Deck struct {
	ID        string
	Shuffled  bool
	Cards     Cards
	DrawCount int
	sync.Mutex
}

var decks = make(map[string]*Deck)

func (c Cards) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
}

func (c Cards) validate() bool {
	for _, card := range c {
		if !slices.Contains(fullDeckCards[:], card) {
			return false
		}
	}
	return true
}
