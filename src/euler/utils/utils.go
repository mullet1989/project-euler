package utils

import (
	"math"
)

// IsPrime ... this is a function
func IsPrime(n int) (isPrime bool) {
	sqrt := int(math.Sqrt(math.Abs(float64(n))))
	for i := 2; i < sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}