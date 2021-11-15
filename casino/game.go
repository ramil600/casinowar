package casino

import (
	"encoding/json"
	"log"
)

type State struct {
	Dealer   Card
	Player   Card
	Cards    Deck
	OrigBet  int
	SideBet  int
	TotalBet int
	War      bool
	TopCard  int
}

func (s *State) PlaceBet(amt int) {
	s.OrigBet = amt
}

func (s *State) GotoWar(war bool) {
	s.War = war
}

func (s *State) PlaceSideBet(amt int) {
	s.SideBet = amt
}

func PlayerWin(s *State, u *User) {
	u.bank += s.TotalBet
}

func InitState(deck Deck) *State {
	return &State{
		Cards: deck,
	}
}

func (s *State) DealCards() TCPData {

	d := (s.Cards)
	s.Player = d[s.TopCard]
	s.Dealer = d[s.TopCard+1]
	s.TopCard = s.TopCard + 2

	cardsDealed := CardsDealed{
		PlayerCard: s.Player,
		DealerCard: s.Dealer,
		OrigBet:    s.OrigBet,
	}
	bytes, err := json.Marshal(cardsDealed)
	if err != nil {
		log.Fatal("game.go: Could not marshal json")
	}

	return TCPData{
		Typ:  "cardsdealed",
		Data: bytes,
	}

}
