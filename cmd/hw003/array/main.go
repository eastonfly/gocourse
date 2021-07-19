package main

import (
	"fmt"

	"github.com/eastonfly/gocourse/cmd/hw003/array/average"
)

func main() {
	fmt.Println("Please, enter the numbers:")
	var line string
	fmt.Scanf("%s", &line)
	ave := average.FindAver(line)

	fmt.Printf("\n %.2f", ave)
}
