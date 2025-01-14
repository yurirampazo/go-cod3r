package math

import (
	"fmt"
	"strconv"
)

func Average(numbers ...float64) float64 {
	total := 0.0

	for _, n := range numbers {
		total += n
	}

	avg := total / float64(len(numbers))

	roundedAvg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return roundedAvg
}