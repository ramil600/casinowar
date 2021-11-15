package casinowar

type State struct {
	Dealer   *Card
	Player   *Card
	Cards    *Deck
	OrigBet  int
	SideBet  int
	TotalBet int
	War      bool
	TopCard  int

}

func (s *State) PlaceBet(amt int) {
	s.OrigBet = amt
}

func(s *State) GotoWar(war bool) {
	s.War = war
}

func (s *State) PlaceSideBet(amt int) {
	s.SideBet = amt
}


func PlayerWin (s *State, u *User) {
	u.bank += s.TotalBet
}

func InitState(deck *Deck) *State {
	return &State{
		Cards: deck,
	}
}

func (s *State) DealCards() {
	s.Player = &s.Cards[s.TopCard]
	s.Dealer = &s.Cards[s.TopCard + 1]
	s.TopCard = s.TopCard + 2


}
