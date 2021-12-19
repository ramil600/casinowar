package casino

import (
	"encoding/json"
	"errors"
)

// State is the current state of the game
type State struct {
	DCard   Card // card of Dealer
	PCard   Card // card of Player
	Cards   Deck // deck of cards
	Player  User // player information: original bet, side bet
	War     bool // if player decided to go to war
	TopCard int  // top card of the deck, returns index of first non-drawn card

}

func (s *State) PlaceBet(amt int) {
	s.Player.OrigBet = amt
	s.Player.Bank = s.Player.Bank - float64(amt)
}

func (s *State) GotoWar(war bool) {
	s.War = war
}

func (s State) IsDraw() bool {
	return s.DCard.Rank == s.PCard.Rank
}

func (s State) HasPlayerWon() bool {
	return s.PCard.Rank > s.DCard.Rank
}

func (s *State) PlaceSideBet(amt int) {
	s.Player.SideBet = amt
	s.Player.Bank = s.Player.Bank - float64(amt)
}
func (s *State) PlaceWarBet() {
	s.Player.Bank -= float64(s.Player.OrigBet)
	s.Player.TotalBet += s.Player.OrigBet * 2
}
func (s *State) ProcessWarOut() {
	s.Player.Bank += (float64)(s.Player.OrigBet / 2)
	s.Player.OrigBet = 0
	s.Player.TotalBet = 0
}

func (s *State) PlayerWon() {
	s.Player.Bank += float64(s.Player.OrigBet * 2)
}

func InitState(deck Deck) *State {
	return &State{
		Cards: deck,
		Player: User{
			Bank: 10000,
		},
	}
}
func (s *State) UpdateUserBank() {
	if s.PCard.Rank > s.DCard.Rank {
		s.Player.Bank += float64(s.Player.OrigBet * 2)
	}

	if (s.PCard.Rank == s.DCard.Rank) && s.War {
		s.Player.Bank += float64(s.Player.OrigBet * 4)
		s.Player.Bank += float64(s.Player.SideBet * 10)
	}
}
func (s *State) DealCards() (TCPData, error) {

	d := s.Cards
	if s.TopCard > 52*NumDecks-2 {
		return TCPData{}, errors.New("Reached the end of the deck")
	}
	s.PCard = d[s.TopCard]
	s.DCard = d[s.TopCard+1]
	s.TopCard = s.TopCard + 2

	s.UpdateUserBank()

	cardsDealed := CardsDealed{
		PlayerCard: s.PCard,
		DealerCard: s.DCard,
		OrigBet:    s.Player.OrigBet,
		UserBank:   s.Player.Bank,
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
