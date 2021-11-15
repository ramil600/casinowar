package main

// StartBet is sent by client to initiate the first bet
type StartBet struct {
	Bet int `json:"bet"`
}

