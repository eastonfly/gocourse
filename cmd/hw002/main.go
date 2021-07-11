package main

import (
	"fmt"

	"github.com/eastonfly/gocourse/cmd/hw002/fibon"
)

func main() {
	var n int
	fmt.Println("Vvedi n")
	fmt.Scan(&n)
	fibon.Printer(n)
}
