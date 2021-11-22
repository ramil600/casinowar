package casino

type User struct {
	OrigBet  int
	SideBet  int
	TotalBet int
	Bank     float64
}

func NewUser() *User {
	return &User{}
}
