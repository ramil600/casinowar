package casino

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestParseCardsDealt(t *testing.T) {

	want:= casino.CardsDealed{
	PlayerCard: casino.Card{
		Suit:0,
		Rank:0,
	},
	DealerCard: casino.Card{
		Suit:0,
		Rank:0,
	},
	OrigBet:0,
	SideBet:2,
	TotalBet:0,
	WarDraw:"",
	}
	data, _ := json.Marshal(want)

	tcpdata := casino.TCPData{
		Typ: "cardsdealed",
		Data: data,
	}

	have, err := casino.ParseCardsDealt(tcpdata)

	if err != nil{
		t.Error("ParseCardsDealt could not parse tcp data", err)
	}
	assert.Equal(t,want,have, "ParseCardsDealt should return this struct")

}