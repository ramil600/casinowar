package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/ramil600/casinowar/casino"
)

//PrintResult prints the cardsdealed message coming from server
func PrintResult(message casino.CardsDealed) {

	fmt.Println("Dealed Cards are:")
	fmt.Println("Player:", message.PlayerCard.Rank)
	fmt.Println("Dealer:", message.DealerCard.Rank)
	if message.PlayerCard.Rank == message.DealerCard.Rank {
		fmt.Println("Draw..Do you want to go to war?:[Y]yes or [N]no:")
	} else if message.PlayerCard.Rank > message.DealerCard.Rank {
		fmt.Println("Congradulations you won!")
	} else {
		fmt.Println("Sorry you lost!")
	}
}

// ParseBetInput asks if you want to play another hand, if answer is yes
// it asks for the bet amount and sends message with input to server
func ParseBetInput(r io.Reader, w io.Writer) {

	var regex = `^[0-9]*$`
	validnum := regexp.MustCompile(regex)

	rd := bufio.NewReader(r)

	for {
		fmt.Print("Do you want a new deal? [Y]yes or [N]no:")

		input, err := rd.ReadString('\n')
		if err != nil {
			log.Println("Could not read input", err)
		}
		input = strings.ToLower(input)
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r") // In case compiling in Windows

		if input == "y" || input == "yes" {
			fmt.Println("Input Your bet:")

			input, err = rd.ReadString('\n')
			if err != nil {
				log.Println("Could not read input", err)
			}
			input = strings.TrimRight(input, "\n")
			input = strings.TrimRight(input, "\r") // In case compiling in Windows

			if validnum.MatchString(input) {
				fmt.Println("You entered valid number")
			}

			break

		} else if input == "n" || input == "no" {
			fmt.Println("Okay..Quitting.")
			break
		} else {
			fmt.Println("Sorry I didn't get you..")
			continue
		}
	}

}

func main() {
	//Take addr:host from input and connect to server
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Usage: go run _client.go addr:host")
	}
	addr := args[1]
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Could not connect to host provided: %s", addr)
	}

	fmt.Println(c.RemoteAddr())
	buf := casino.TCPData{}
	dec := json.NewDecoder(c)

	// Accept dealed cards inside the message: type CardsDealed
	if err := dec.Decode(&buf); err != nil {
		log.Println(err)
	}
	cardsdealed, err := casino.ParseCardsDealed(buf)
	if err != nil {
		log.Println(err)
	}

	PrintResult(cardsdealed)
	ParseBetInput(os.Stdin, c)

	bytes, err := json.Marshal(cardsdealed)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(bytes)

	c.Write(bytes)

	//dec := json.NewDecoder(c)
	//
	//var tcpMsg = casino.TCPData{}
	//dec.More()
	//err = dec.Decode(&tcpMsg)
	//if err != nil {
	//	log.Fatal("Could not decode tcp Message:", err)
	//}
	//switch tcpMsg.Typ {
	//case "cardsdealed":
	//	fmt.Println("You received Cards Dealed Struct")
	//default:
	//	fmt.Println("Unknown Message Type")
	//
	//}
	select {}

}
