package casino

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDealCards(t *testing.T) {

	s := State{
		DCard: Card{
			Suit: Clubs,
			Rank: Ace,
		},
		PCard: Card{
			Suit: Diamonds,
			Rank: Ace,
		},
		Cards: *(NewDeck()),
		Player: User{
			Bank:    1000,
			OrigBet: 100,
			SideBet: 10,
		},
		TopCard: 2,
		War:     false,
	}
	have, err := s.DealCards()
	if err != nil {
		t.Error("Test failed when dealing cards", err)
	}

	cardsDealt := CardsDealed{
		PlayerCard: Card{
			Rank: Four,
			Suit: Clubs,
		},
		DealerCard: Card{
			Rank: Five,
			Suit: Clubs,
		},
		OrigBet:  s.Player.OrigBet,
		UserBank: s.Player.Bank,
	}

	data, _ := json.Marshal(cardsDealt)

	want := TCPData{
		Typ:  "cardsdealed",
		Data: data,
	}
	assert.Equal(t, have, want, "Verify if DealCards returns expected msg")

}
