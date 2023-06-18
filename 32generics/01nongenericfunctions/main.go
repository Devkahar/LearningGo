package main

import "fmt"

func SumInts(m map[string]int64) (sum int64) {
	sum = 0
	for _, v := range m {
		sum += v
	}
	return
}

func SumFlots(m map[string]float64) (sum float64) {
	sum = 0
	for _, v := range m {
		sum += v
	}
	return
}

func main() {
	ints := map[string]int64{
		"first":  12,
		"second": 24,
	}
	flots := map[string]float64{
		"first":  22.5,
		"second": 99.6,
	}

	fmt.Printf("Sums with non generic functions %v and %v\n", SumInts(ints), SumFlots(flots))
}
