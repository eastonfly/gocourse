package fibon

import "fmt"

func Printer(n int) {
	switch {
	case n == 0:
		fmt.Println(0)
	case n == 1:
		fmt.Println(0)
	case n < 0:
		fmt.Println("Error: Negative value")
	default:
		var prelast, last uint64 = 0, 1
		var sum uint64
		fmt.Print(0, " ", 1, " ")
		for i := 2; i < n; i++ {
			sum = prelast + last
			fmt.Print(sum, " ")
			prelast = last
			last = sum
		}
	}
}
