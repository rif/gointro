package uuid

import "errors"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("invalid argument")
	}
	prod := 1
	for i := 1; i <= n; i++ {
		prod *= i
	}
	return prod, nil
}
