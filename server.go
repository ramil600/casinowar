package main

import (
	//"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/ramil600/casinowar/casino"
	"net"

	//"net"
	"os"
	"time"
)

// HandleConnection will response to connected client. Server will start HandleConnection as
// a separate goroutine. Context to be implemented.
func HandleConnection(ctx context.Context, c io.ReadWriteCloser) {


	defer c.Close()

	cardsdealed := casino.CardsDealed{}

	//Initialize new deck
	cards := casino.NewDeck()
	cards.Shuffle(20)
	state := casino.InitState(*cards)

	//Deal Cards
	msg := state.DealCards()
	rawMessage, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Could not marshal message")
	}
	dec := json.NewDecoder(c)


	// Send message to player
	if _, err := c.Write(rawMessage); err != nil {
		log.Fatal(err, ": Could not write to c")
	}
	time.Sleep(500 * time.Millisecond)

	// Receive the input from the user
	dec.Decode(&cardsdealed)
	fmt.Println(cardsdealed)

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
