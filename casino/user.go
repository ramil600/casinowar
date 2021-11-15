package casino

type User struct {
	bank int
}


func NewUser() *User {
	return &User{}
}

func (u *User) IncrementBank(win int) {
	u.bank = u.bank + win
}
func (u User) GetBank() int {
	return u.bank
}
