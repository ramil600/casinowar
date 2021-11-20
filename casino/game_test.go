package casino

import(
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DealCards(t * testing.T){

	s := casino.State{
		DCard: casino.Card{
			Suit: casino.Diamonds,
			Rank: casino.Ace,
		},
		PCard: casino.Card{
			Suit: casino.Diamonds,
			Rank: casino.Ace,
		},
		Cards: *(casino.NewDeck()),
		Player: casino.User{
			Bank: 1000,
			OrigBet: 100,
			SideBet:10,
		},
		TopCard:2,
		War: false,
	}
	have, err:= s.DealCards()
	if err != nil {
		t.Error("Test failed when dealing cards", err)
	}


	cardsDealt := casino.CardsDealed{
		PlayerCard: casino.Card{
			Rank: casino.Four,
			Suit: casino.Clubs,
		},
		DealerCard: casino.Card{
			Rank: casino.Five,
			Suit: casino.Clubs,
		},
		OrigBet: s.Player.OrigBet,
	}

	data, _ := json.Marshal(cardsDealt)

	want := casino.TCPData{
		Typ: "cardsdealed",
		Data : data,
	}
	assert.Equal(t, have, want, "Verify if DealCards returns expected msg")

}
