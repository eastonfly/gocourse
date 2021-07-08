package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	message := emoji.Sprint("eat :sandwich: sleep :zzz: Golang :technologist: ")
	fmt.Println(message)
}
