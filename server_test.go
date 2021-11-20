package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/ramil600/casinowar/casino"
	"github.com/stretchr/testify/assert"
)


/*
type Mockcon interface {

	io.Reader
	io.Writer
}

type mockon struct{
	*strings.Reader
	*bufio.Writer
}

func (m mockon) Close() error  {
	return nil
}

func  TestHandle_Connection(t *testing.T)  {

	var bufw bytes.Buffer

	newBet := casino.StartBet{
		Bet: 10,
		SideBet: 10,
		WarReq: "false",
	}


	output := bufio.NewWriter(&bufw)

	msg,_ := json.Marshal(newBet)

	input:= strings.NewReader(string(msg))

	mockc := mockon{
		Reader: input,
		Writer: output,
	}

	ctx, cancel := context.WithCancel(context.Background())
	HandleConnection(ctx,mockc)
	cancel()
	input = nil
	time.Sleep( time.Millisecond)

	have := bufw.Bytes()
	fmt.Println(string(have))




	assert.Equal(t, "TRUE", have, "Testing HandleConnection")
*/
}