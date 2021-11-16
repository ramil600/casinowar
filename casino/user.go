package casino

type User struct {
	OrigBet  int
	SideBet  int
	TotalBet int
	Bank     int
}

func NewUser() *User {
	return &User{}
}

func (u *User) IncrementBank(win int) {
	u.Bank = u.Bank + win
}

func (u User) GetBank() int {
	return u.Bank
}
