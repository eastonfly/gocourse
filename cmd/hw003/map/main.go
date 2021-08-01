package main

import (
	"fmt"
	"sort"
)

func printSorted(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("Sorted map: ")
	for _, k := range keys {
		fmt.Printf("%d:%s ", k, m[k])
	}
}

func main() {
	m := map[int]string{
		9:  "hello",
		2:  "a",
		0:  "b",
		1:  "c",
		99: "ololo",
	}
	fmt.Println("Original map:", m)
	printSorted(m)

}
