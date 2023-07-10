package main

import "fmt"

// Fibonacci returns the n-th Fibonacci number.
// Fibonacci number cannot be represented as a uint64.
func Fibonacci(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}

	// Error handling for fibonacci uint
	if n > 93 {
		return 0, fmt.Errorf("unsupported Fibonacci number %d: too large", n)
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}
