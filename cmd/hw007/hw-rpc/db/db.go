package db

import (
	"errors"
	"fmt"
)

var Base = []string{
	1:  "Hey",
	2:  "I",
	3:  "Just",
	4:  "Met",
	5:  "You",
	6:  "And",
	7:  "This",
	8:  "Is",
	9:  "Crazy",
	10: "But",
	11: "Here's",
	12: "My",
	13: "Number",
	14: "So",
	15: "Call",
	16: "Me",
	17: "Maybe",
}

func Create(name string) int {
	l := len(Base)
	Base = append(Base, name)
	return l
}

func Show() {
	for key := 1; key < len(Base); key++ {
		fmt.Printf("%d : %s\n", key, Base[key])
	}
}

func Check(name string) error {
	for _, val := range Base {
		if name == val {
			return errors.New("such user is already exist")
		}
	}
	return nil
}
