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
	defer fmt.Println("Here it is the row of",n,"Fibonachi numbers")
	defer fmt.Println("\n----------------------------")
	fibon.Printer(n)
}
