package casinowar

// StartBet is sent by client to initiate the first bet
type StartBet struct {
	Bet int `json:"bet"`
}
// CardsDealed sent by server with initiated bet and dealed hand
type CardsDealed struct {
	PlayerCard Card
	DealerCard Card
	OrigBet int `json:"orig_bet"`
	SideBet int `json:"side_bet"`
	TotalBet int `json:"total_bet"`
}

type WarRequest struct {
	WarReq string `json:"war_req"`
}
type TCPData struct {
	Typ string
	Data []byte
}