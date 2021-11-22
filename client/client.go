package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ardanlabs/conf"
	"github.com/ramil600/casinowar/casino"
)

//PrintCardsDealt prints the cardsdealed message coming from server
func PrintCardsDealt(message casino.CardsDealed) {

	var mranks = map[casino.Ranks]string{casino.Two: "Two", casino.Three: "Three",
		casino.Four: "Four", casino.Five: "Five", casino.Six: "Six", casino.Seven: "Seven",
		casino.Eight: "Eight", casino.Nine: "Nine", casino.Ten: "Ten", casino.Jack: "Jack",
		casino.Queen: "Queen", casino.King: "King", casino.Ace: "Ace"}

	var msuits = map[casino.Suits]string{casino.Clubs: "Clubs", casino.Diamonds: "Diamonds",
		casino.Hearts: "Hearts", casino.Spades: "Spades"}

	fmt.Println("Dealt Cards are:")
	fmt.Println("Player:", mranks[message.PlayerCard.Rank], "of", msuits[message.DealerCard.Suit])
	fmt.Println("Dealer:", mranks[message.DealerCard.Rank], "of", msuits[message.DealerCard.Suit])
	if message.PlayerCard.Rank == message.DealerCard.Rank {
		fmt.Println("You are tied..")
	} else if message.PlayerCard.Rank > message.DealerCard.Rank {
		fmt.Println("Congratulations you won!")
	} else {
		fmt.Println("Sorry you lost!")
	}
	fmt.Printf("You have $%.2f CAD left in your Bank.\n", message.UserBank)
}

//SendWarRequest will query user if he wants war and then enter War Game loop
//Place side bet is optional. Cards are dealt back by server and win/lose calculated
func SendWarRequest(ctx context.Context, r io.Reader, w io.Writer) int {

	rd := bufio.NewReader(r)

	for {
		fmt.Println("Do you want a to go to War?")
		fmt.Println("Your original bet will be doubled")
		fmt.Println("If you win you only win your original bet, and if lose you lose all your bet:[Y/y]yes or [N/n]no")
		input, err := rd.ReadString('\n')
		if err != nil {
			log.Println("Could not read input", err)
		}
		input = strings.ToLower(input)
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r") // In case compiling in Windows

		if input == "y" || input == "yes" {
			//send war request message and start war game scenario
			ParseWarBet(ctx, "true", w)
			return 0
		} else if input == "n" || input == "no" {
			ParseWarBet(ctx, "false", w)
			return -1
		} else {
			fmt.Println("Sorry I didn't get you..")
			continue
		}
	}
}

// ParseWarBet will use io.Reader to parse side bet from user and send a message to tcp connection
func ParseWarBet(ctx context.Context, ans string, w io.Writer) {

	warRequest := casino.StartBet{
		WarReq: ans,
	}
	betMsg, err := json.Marshal(warRequest)
	if err != nil {
		log.Fatal("client.go: Could not marshal bet message", err)
	}
	w.Write(betMsg)
}

// ParseOrigBet will use bufio.Reader to parse initial bet from user and send a message to tcp connection
func ParseOrigBet(ctx context.Context, r io.Reader, w io.Writer) {

	var regex = `^[0-9]*$`
	validNum := regexp.MustCompile(regex)
	var origBet, sideBet int
	rd := bufio.NewReader(r)

	for {
		//Input your bet from os.Stdin
		fmt.Println("Input Your bet:")
		input, err := rd.ReadString('\n')
		if err != nil {
			log.Println("Could not read input", err)
		}
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r") // In case compiling in Windows

		//Validate your bet and check if >0
		if validNum.MatchString(input) {
			origBet, err = strconv.Atoi(input)
			if err != nil {
				log.Println("Could not convert input to number", err)
			}
			if origBet > 0 {
				break
			} else {
				fmt.Println("Please input the positive bet")
			}

		} else {
			fmt.Println("Please input the valid number:")
		}
	}

	// Parse Side Bet, allowed to be 0
	for {
		fmt.Println("Input Your Side Bet. If you wish not to you can simply put 0:")
		input, err := rd.ReadString('\n')
		if err != nil {
			log.Println("Could not read input", err)
		}
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r") // In case compiling in Windows

		if validNum.MatchString(input) {
			fmt.Println("Your bet is accepted.")
			sideBet, err = strconv.Atoi(input)
			if err != nil {
				log.Println("Could not convert input to number", err)
			}
			break
		} else {
			fmt.Println("Please input the valid number")
		}
	}

	newBet := casino.StartBet{
		Bet:     origBet,
		SideBet: sideBet,
	}
	betMsg, err := json.Marshal(newBet)
	if err != nil {
		log.Fatal("client.go: Could not marshal bet message", err)
	}
	w.Write(betMsg)
}

// ParseBetInput asks if you want to play another hand, if answer is yes
// it asks for the bet amount and sends message with input to server
func ParseBetInput(ctx context.Context, r io.Reader, w io.Writer) int {

	//ctx, cancel := context.WithCancel(ctx)
	rd := bufio.NewReader(r)
	for {
		fmt.Print("Do you want a new deal? [Y/y]yes or [N/n]no:")
		input, err := rd.ReadString('\n')
		if err != nil {
			log.Println("Could not read input", err)
		}
		input = strings.ToLower(input)
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r") // In case compiling in Windows

		if input == "y" || input == "yes" {
			ParseOrigBet(ctx, os.Stdin, w)
			return 0
		} else if input == "n" || input == "no" {
			fmt.Println("Okay..Quitting.")
			return -1
		} else {
			fmt.Println("Sorry I didn't get you..")
			continue
		}
	}
}

func main() {

	// SETTING UP CLI CONFIGURATION
	var cfg struct {
		Host string `conf:"default:0.0.0.0"`
		Port string `conf:"default:8081"`
	}

	if err := conf.Parse(os.Args[1:], "SERVER", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("SERVER", &cfg)
			if err != nil {
				log.Fatal("generating config usage")
			}
			fmt.Println(usage)

		}
		log.Fatal("Parsing config")

	}

	// Connect to the game server
	addr := cfg.Host + ":" + cfg.Port
	c, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalf("Could not connect to host provided: %s", addr)
	}
	defer c.Close()

	fmt.Println("You are connected to a remote game server:", c.RemoteAddr())
	buf := casino.TCPData{}
	dec := json.NewDecoder(c)
	ctx := context.Background()

	for {
		//Accept Initial Bet from the player, if he wants to quit exit the game loop
		if ParseBetInput(ctx, os.Stdin, c) == -1 {
			break
		}
		// Accept dealt cards inside the message: type CardsDealed
		if err := dec.Decode(&buf); err != nil {
			log.Println(err)
			break
		}
		cardsDealt, err := casino.ParseCardsDealt(buf)
		if err != nil {
			log.Println(err)
		}
		_, err = json.Marshal(cardsDealt)
		if err != nil {
			log.Println(err)
		}
		PrintCardsDealt(cardsDealt)

		if cardsDealt.WarDraw == "true" {
			if SendWarRequest(ctx, os.Stdin, c) == -1 {
				continue
			}

			// Accept dealt cards inside the message: type CardsDealed
			if err := dec.Decode(&buf); err != nil {
				log.Println(err)
				break
			}
			cardsDealt, err := casino.ParseCardsDealt(buf)
			if err != nil {
				log.Println(err)
			}
			_, err = json.Marshal(cardsDealt)
			if err != nil {
				log.Println(err)
			}
			PrintCardsDealt(cardsDealt)
		}
	}
	fmt.Println("Thank you for playing, hope to see you again!")
}
