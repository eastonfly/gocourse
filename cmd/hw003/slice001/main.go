package main

import (
	"fmt"
	"strings"
)

func max(slice []string) string {
	var maxim string = slice[0]
	for _, val := range slice {
		if len(val) > len(maxim) {
			maxim = val
		}
	}
	return maxim
}

func main() {
	var line string
	fmt.Println("Please, enter the slice of strings:")
	fmt.Scanf("%s", &line)
	sli := strings.Split(line, ",")
	fmt.Printf("%s", max(sli))
}
