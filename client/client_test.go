package main

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ramil600/casinowar/casino"
	"github.com/stretchr/testify/assert"
)


func TestParseWarBet( t *testing.T){

	var b bytes.Buffer
	msg := casino.StartBet{
		WarReq: "true",
	}
	want, _ := json.Marshal(msg)

	ParseWarBet(context.TODO(), "true",&b)
	have := b.Bytes()

	assert.Equal(t, have, want, "ParseWarBet message is tested")

}

func TestParseOrigBet(t *testing.T){

	var r = strings.NewReader("10\n10\n")
	var b bytes.Buffer

	msg := casino.StartBet{
		Bet: 10,
		SideBet: 10,
	}
	want, _ := json.Marshal(msg)
	ParseOrigBet(context.TODO(), r, &b)
	have := b.Bytes()

	assert.Equal(t, want, have, "Checking Whether ParseOrigBet creates msg as required" )

}