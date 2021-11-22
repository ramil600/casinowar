package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ramil600/casinowar/casino"
	"github.com/stretchr/testify/assert"
)

func TestSendCards(t *testing.T) {
	var bufw bytes.Buffer

	tcpMsg := casino.TCPData{
		Typ:  "cardsdealed",
		Data: []byte("some data"),
	}

	want, _ := json.Marshal(tcpMsg)

	SendCards(tcpMsg, &bufw)
	have := bufw.Bytes()
	assert.Equal(t, want, have, "Checking SendCards is sending json to connection")
}

