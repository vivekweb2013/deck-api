package httpservice

type CardResponse struct {
	Value string `json:"value"`
	Suite string `json:"suite"`
	Code  string `json:"code"`
}

type DeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}
