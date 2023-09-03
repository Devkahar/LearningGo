package calculator

import "errors"

func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("Cannot Add Negative value.")
	}
	return a + b, nil
}

func Subtract(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func Mod(a, b int) int {
	return a % b
}
