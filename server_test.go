package main

import (
	"bufio"
	"bytes"
	//"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/ramil600/casinowar/casino"
)



type Mockcon interface {

	io.ReadWriter
}

type mockon struct{
	*bufio.ReadWriter
}
func (m mockon) Close() error  {
	return nil
}

func  TestHandle_Connection(t *testing.T)  {
	buf := make([]byte,1024)
	var bufw bytes.Buffer

	newBet := casino.StartBet{
		Bet: 10,
		SideBet: 10,
		WarReq: "false",
	}

	input:= bufio.NewReader(strings.NewReader("Hello World!\n"))

	msg,_ := json.Marshal(newBet)

	fmt.Println(string(msg))


	mockc := mockon{
		ReadWriter: bufio.NewReadWriter(input,bufio.NewWriter(&bufw)),
	}
	//go HandleConnection(context.TODO(),mockc)

	if _,err:= mockc.Write(msg); err != nil{
		t.Error(err)
	}
	bufw.Read(buf)
	fmt.Println(string(buf))
	n, err:=  mockc.Read(buf)

	if err != nil {
		t.Error(n, err)
	}

	fmt.Println(string(buf))

	if len(buf) == 0 {
		t.Errorf("The length is 0")
	}
}