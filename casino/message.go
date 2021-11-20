package casino

//StartBet is sent by client to initiate the first original and side bet
type StartBet struct {
	Bet int `json:"bet"`
	SideBet int `json:"side_bet"`
	WarReq string `json:"war_req"`
}
// CardsDealed sent by server with initiated bet and dealed hand
type CardsDealed struct {
	PlayerCard Card
	DealerCard Card
	OrigBet int `json:"orig_bet"`
	SideBet int `json:"side_bet"`
	TotalBet int `json:"total_bet"`
	UserBank int `json:"user_bank"`
	WarDraw string `json:"war_draw"`
}

type WarRequest struct {
	WarReq string `json:"war_req"`
	SideBet int `json:"side_bet"`
}
type TCPData struct {
	Typ string
	Data []byte
}