package average

import (
	"log"
	"strconv"
	"strings"
)

func FindAver(line string) float64 {
	var (
		sum, count int
	)
	ar := strings.Split(line, ",")
	for i, val := range ar {
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + n
		count = i
	}
	return float64(sum / count)
}
