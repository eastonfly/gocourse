package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Users struct {
	Name string
}

func main() {

	serverAddress := "127.0.0.1"

	client, err := rpc.DialHTTP("tcp", serverAddress+":8081")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Println("Usage:")
	fmt.Println("1. To add new user type: register username")
	fmt.Println("2. To view existing users type: list")

	in := bufio.NewReader(os.Stdin)
	choise, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("input error: ", err)
	}

	mes := strings.Split(choise, " ")

	if len(mes) != 1 {
		mes[0] = mes[0] + "\n"
	}
	switch mes[0] {

	case "register\n":
		{

			if len(mes) != 2 {
				log.Fatal("username should not be empty")
			}
			user := Users{mes[1]}
			var reply int
			err = client.Call("Handler.Register", user, &reply)
			if err != nil {
				log.Fatal("register error: ", err)
			}
			fmt.Printf("User with name: %s and id: %d sucessfully created\n", user.Name, reply)
		}

	case "list\n":
		{
			var (
				list []string
				use  = Users{}
			)
			err = client.Call("Handler.List", use, &list)
			if err != nil {
				log.Fatal("show error: ", err)
			}
			for key := 1; key < len(list); key++ {
				fmt.Printf("%d : %s\n", key, list[key])
			}
		}
	default:
		fmt.Printf("Not supported!\n")
	}
}
