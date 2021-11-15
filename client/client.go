package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"log"

	"net"
	"os"

	"github.com/ramil600/casinowar/casino"
)

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

	//reader := bufio.NewReader(os.Stdin)
	fmt.Println(c.RemoteAddr())
	buf := casino.TCPData{}
	dec := json.NewDecoder(c)

	if err := dec.Decode(&buf); err != nil {
		log.Println(err)
	}
	cardsdealed, err := casino.ParseCardsDealed(buf)

	if err != nil{
		log.Println(err)
	}

	bytes, err := json.Marshal(cardsdealed)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(bytes)

	c.Write(bytes)

	//go func () {
	//	for server.Scan() {
	//		buf := server.Bytes()
	//
	//		if err != nil {
	//			log.Fatal("Could not get the message back from server")
	//		}
	//
	//		fmt.Println("Server replied: ", string(buf))
	//	}
	//}()

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

	//for {
	//	fmt.Print(">>")
	//	// Read text from stdin
	//	text,err := reader.ReadString('\n')
	//	if err != nil {
	//		log.Fatal("Error encountered reading from connection")
	//	}
	//	// Write text to server
	//	fmt.Fprint(c, text)
	//
	//	srvtext, err:= server.ReadString('\n')
	//	if err != nil {
	//		log.Fatal("Could not get the message back from server")
	//	}
	//	fmt.Println("Server replied: ", srvtext)
	//
	//
	//}
}
