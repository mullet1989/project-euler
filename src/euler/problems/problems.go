package problems

import (
	"fmt"
	"math"
)

func Problem23() {
	nums := getAbundantNumbers(28123)
	sums := doCombinations(nums)

	var sum int
	for i := 1; i <= 28123; i++ {
		if _, ok := sums[i]; !ok {
			sum += i
		}
	}

	fmt.Println(sum)
}

func getDivisors(n int) (sum int) {
	squareRoot := int(math.Sqrt(float64(n)))
	var sumLocal = 1
	for i := 2; i <= squareRoot; i++ {
		if n%i == 0 {
			sumLocal += i
			if i != n/i {
				sumLocal += n / i // paired factors
			}
		}
	}
	return sumLocal
}

func getAbundantNumbers(upperBound int) (nums []int) {
	var abundantNums []int
	for i := 1; i <= upperBound; i++ {
		sum := getDivisors(i)
		if sum > i {
			abundantNums = append(abundantNums, i)
		}
	}
	return abundantNums
}

func doCombinations(abundantNums []int) (sums map[int]int) {
	m := make(map[int]int)
	for index := range abundantNums {
		var i = index
		for i < len(abundantNums)-index {
			sum := abundantNums[index] + abundantNums[i]
			m[sum] = sum
			i++
		}
	}
	return m
}
