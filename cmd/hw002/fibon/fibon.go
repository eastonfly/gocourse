package fibon

import "fmt"

func Printer(n int) {
	switch {
	case n == 0:
		fmt.Println(0)
	case n == 1:
		fmt.Println(1)
	case n < 0:
		fmt.Println("Error: Negative value")
	default:
		var prelast, last uint64 = 1, 1
		var sum uint64
		fmt.Print(1, " ", 1, " ")
		for i := 2; i < n; i++ {
			sum = prelast + last
			fmt.Print(sum, " ")
			prelast = last
			last = sum
		}
	}
}
