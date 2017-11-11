package main

import (
	problems "euler/problems"
	permute "euler/permute"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	problems.Problem23()
	fmt.Println(permute.Reverse("Hello World"))
}

func p23() {
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

func p26(n int) int {
	sequenceLength := 0

	for i := n; i > 1; i-- {
		if sequenceLength >= i {
			break
		}

		foundRemainders := make([]int, i)
		value, position := 1, 0

		for foundRemainders[value] == 0 && value != 0 {
			foundRemainders[value] = position
			value *= 10
			value %= i
			position++
		}

		if position-foundRemainders[value] > sequenceLength {
			sequenceLength = position - foundRemainders[value]
		}
	}
	return sequenceLength + 1
}

func p28(n int) int {
	init := 1
	for i := 3; i <= n; i += 2 {
		power := int(math.Pow(float64(i), 2))
		init += (power + (power - (i - 1)) + (power - 2*(i-1)) + (power - 3*(i-1)))
	}
	return init
}

func p29(uppera, upperb int) map[string]int {
	values := make(map[string]int)
	for i := 2; i <= uppera; i++ {
		for j := 2; j <= upperb; j++ {
			biga := big.NewInt(int64(i))
			bigb := big.NewInt(int64(j))
			val := biga.Exp(biga, bigb, nil)
			values[val.String()] = 0
		}
	}
	return values
}

func p30() {
	var i int64 = 2
	var total int64 = 0
	for i < 10000000 {
		stringi := strconv.FormatInt(i, 10)
		var accum float64 = 0
		for _, j := range stringi {
			c := string(j)
			c1, _ := strconv.ParseFloat(c, 64)
			accum += math.Pow(c1, 5)
		}
		if int64(accum) == i {
			fmt.Println(accum)
			total += int64(accum)
		}
		i++
	}
	fmt.Println(total)

}

func binarySearch(array []int, start, end int) int {
	mid := start + ((end - start) / 2)

	if mid == start {
		if end == len(array)-1 {
			return len(array)
		}
		return array[start] + 1
	}

	if array[mid] != mid {
		return binarySearch(array, start, mid)
	}
	return binarySearch(array, mid, end)
}
