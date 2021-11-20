package casino

import (
	"encoding/json"
)

// Current state of the game
type State struct {
	DCard   Card // card of Dealer
	PCard   Card // card of Player
	Cards   Deck // deck of cards
	Player  User // player information: original bet, side bet
	War     bool // if player decided to go to war
	TopCard int  // top card of the deck, returns index of first non drawn card
}

func (s *State) PlaceBet(amt int) {
	s.Player.OrigBet = amt
	s.Player.Bank = s.Player.Bank - amt
}

func (s *State) GotoWar(war bool) {
	s.War = war
}

func (s State) IsDraw() bool{
	return s.DCard.Rank == s.PCard.Rank
}

func (s *State) PlaceSideBet(amt int) {
	s.Player.SideBet = amt
}

func PlayerWin(s *State, u *User) {
	u.Bank += s.Player.TotalBet
}

func InitState(deck Deck) *State {
	return &State{
		Cards: deck,
	}
}

func (s *State) DealCards() (TCPData, error) {

	d := (s.Cards)
	s.PCard = d[s.TopCard]
	s.DCard = d[s.TopCard+1]
	s.TopCard = s.TopCard + 2

	cardsDealed := CardsDealed{
		PlayerCard: s.PCard,
		DealerCard: s.DCard,
		OrigBet:    s.Player.OrigBet,

	}
	if cardsDealed.PlayerCard.Rank == cardsDealed.DealerCard.Rank {
		cardsDealed.WarDraw = "true"
	}

	bytes, err := json.Marshal(cardsDealed)
	if err != nil {
		return TCPData{}, err
	}

	return TCPData{
		Typ:  "cardsdealed",
		Data: bytes,
	}, nil

}
