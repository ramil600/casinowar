package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		log.Fatal("Usage: go run _client.go addr:host")
	}

	addr := args[1]

	c, err := net.Dial("tcp",addr)
	if err != nil {
		log.Fatalf("Could not connect to host provided: %s", addr)
	}

	reader := bufio.NewReader(os.Stdin)
	server := bufio.NewReader(c)

	for {
		fmt.Print(">>")
		// Read text from stdin
		text,err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error encountered reading from connection")
		}
		// Write text to server
		fmt.Fprint(c, text)

		srvtext, err:= server.ReadString('\n')
		if err != nil {
			log.Fatal("Could not get the message back from server")
		}
		fmt.Println("Server replied: ", srvtext)


	}
}
