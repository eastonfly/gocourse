package main

import (
	"fmt"

	"github.com/eastonfly/gocourse/cmd/hw002/fibon"
)

func main() {
	var n int
	fmt.Println("Please, enter the amount:")
	fmt.Scan(&n)
	fmt.Println("------------------------------")
	defer fmt.Printf("Here it is the row of %d Fibonachi numbers", n)
	defer fmt.Println("\n----------------------------")
	fibon.Printer(n)
}
