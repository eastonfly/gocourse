package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	// Set port listening
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	// Open port
	conn, err := ln.Accept()
	if err != nil {
		log.Print(err)
	}

	// Start loop
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//Try to conver str to int
		i, error := strconv.Atoi(message[:(len(message) - 1)])
		if error != nil {
			newmessage := strings.ToUpper(message)
			conn.Write([]byte(newmessage + "\n"))
			fmt.Println("It's string, UPPERCASE it")
		} else {
			mes := strconv.FormatInt(int64(i*2), 10)
			conn.Write([]byte(mes + "\n"))
			fmt.Println("Int multiple by 2")

		}

	}
}
