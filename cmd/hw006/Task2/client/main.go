package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	return text
}

func main() {

	// Connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		text := Read()
		if text == "exit\n" {
			fmt.Println("Exiting due \"Exit\" was entered!")
			break
		}
		// Send to server
		fmt.Fprintf(conn, text+"\n")
		// Recieve the answer
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
