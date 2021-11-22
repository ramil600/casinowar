package casino

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCardsDealt(t *testing.T) {

	want := CardsDealed{
		PlayerCard: Card{
			Suit: 0,
			Rank: 0,
		},
		DealerCard: Card{
			Suit: 0,
			Rank: 0,
		},
		OrigBet:  0,
		SideBet:  2,
		TotalBet: 0,
		WarDraw:  "",
	}
	data, _ := json.Marshal(want)

	tcpdata := TCPData{
		Typ:  "cardsdealed",
		Data: data,
	}

	have, err := ParseCardsDealt(tcpdata)

	if err != nil {
		t.Error("ParseCardsDealt could not parse tcp data", err)
	}
	assert.Equal(t, want, have, "ParseCardsDealt should return this struct")

}
