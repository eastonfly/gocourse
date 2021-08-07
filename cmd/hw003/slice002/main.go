package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func reverse(slice []int64) []int64 {
	var slice2 = make([]int64, len(slice), cap(slice))
	for i := (len(slice) - 1); i >= 0; i-- {
		slice2[(len(slice)-1)-i] = slice[i]
	}
	return slice2
}

func main() {
	fmt.Println("Please, enter the slice:")
	var (
		line string
		sli  []int64
	)
	fmt.Scanf("%s", &line)
	ar := strings.Split(line, ",")
	for _, val := range ar {
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sli = append(sli, int64(n))
	}
	fmt.Printf("\nOriginal slice: %v", sli)
	fmt.Printf("\nReverse slice: %v", reverse(sli))

}
