package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"

	//"net"
	"os"
)

func HandleConnection (ctx context.Context, c net.Conn){

	user := NewUser()
	cards := NewDeck()
	cards.Shuffle(20)
	state := InitState(cards)

	for {
		state.DealCards()
		r := bufio.NewReader(c)
		text, err := r.ReadString('\n')

		if err != nil {
			log.Println("Received message from client", err)
			break
		}
		fmt.Println(text)

		fmt.Fprintf(c, "Your Card: %v ", state.Player.Suit)
		fmt.Fprintf(c, "Dealer Card %v\n", state.Dealer.Suit)

		user.GetBank()

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
		go HandleConnection(context.Background(),c)

	}

}
