package main

import "fmt"

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) (sum V) {
	sum = 0
	for _, v := range m {
		sum += v
	}
	return
}

// 1. Declare a SumIntsOrFloats function with two type parameters
// (inside the square brackets), K and V, and one argument that uses
// the type parameters, m of type map[K]V. The function returns a value of type V.
// 2. Specify for the K type parameter the type constraint comparable.
// Intended specifically for cases like these, the comparable constraint is predeclared
// in Go. It allows any type whose values may be used as an operand of the comparison
// operators == and !=. Go requires that map keys be comparable.
// So declaring K as comparable is necessary so you can use K as the key in
// the map variable. It also ensures that calling code uses an allowable type for map keys.
// 3. Specify for the V type parameter a constraint that is a union of two types:
// int64 and float64. Using | specifies a union of the two types, meaning that this
//  constraint allows either type. Either type will be permitted by the compiler as an
//  argument in the calling code.
// 4. Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters.
// Note that we know map[K]V is a valid map type because K is a comparable type.
// If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V.

// Declaring a type constraint.

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) (sum V) {
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

	fmt.Printf("Sums with generic functions %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(flots))
	// We need to specify type argument in function call when function has no argument.
	// like this SumIntsOrFloats[string,float64](floats)

	// Using type constraint
	fmt.Printf("Sum Numbers with generic functions %v and %v\n", SumNumbers(ints), SumNumbers(flots))

}
