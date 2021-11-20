package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/ramil600/casinowar/casino"
)

func SendCards (msg casino.TCPData, w io.Writer) error{

	rawMessage, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	
	// Send message to player
	if _, err := w.Write(rawMessage); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)
	return nil
}

// HandleConnection will response to connected client. Server will start HandleConnection as
// a separate goroutine. Context to be implemented.


func HandleConnection(ctx context.Context, c io.ReadWriteCloser) {


	defer c.Close()

	newBet := casino.StartBet{}
	warBet := casino.WarRequest{}

	//Initialize new deck
	cards := casino.NewDeck()
	cards.Shuffle(20)
	state := casino.InitState(*cards)
	dec := json.NewDecoder(c)


	for {
		// Receive the input from the user
		dec.Decode(&newBet)
		fmt.Println(newBet)

		state.PlaceBet(newBet.Bet)
		state.PlaceSideBet(newBet.SideBet)

		if (newBet.Bet <= 0) && (newBet.WarReq != "true") {
			break
		}

		if newBet.WarReq == "true" {
			fmt.Println("War request from player processed.")
		}

		//Deal Cards and send player cardsdealed message
		msg, err := state.DealCards()
		if err != nil {
			log.Fatal("Error during dealing cards generating TCPdata",err)
		}
		if err:= SendCards(msg, c); err != nil {
			log.Println("Could not send the message to player:", err)
			return

		}


		if state.IsDraw(){
			dec.Decode(&warBet)
			if warBet.WarReq == "true"{
				fmt.Println("We are going to war..")
				//state.BurnCards()
				//Deal Cards and Send message to the player
				state.GotoWar(true)
				msg, err := state.DealCards()
				if err != nil {
					log.Fatal("Error during Dealing Cards in war bet", err)
				}
				if err := SendCards(msg, c); err != nil {
					log.Println("Could not send the message to player", err)
					return

				}

			}
			continue
		}
	}





	select {
	case err := <-ctx.Done():
		c.Close()
		log.Println(err)
	}

}

func main() {

	args := os.Args
	if len(args) == 1 {

		log.Fatalln("Please provide host:port")
	}
	// Usage: First argument will be the port to listen at
	port := args[1]

	l, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatalf("Could not open connection to host: %s ", port)
	}
	defer l.Close()

	for {
		// Accept the connection from a client
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Could not accept the connection")
		}
		log.Printf("You accepted connection from: %s", c.RemoteAddr())

		// Handle the received connection
		go HandleConnection(context.Background(), c)

	}

}
