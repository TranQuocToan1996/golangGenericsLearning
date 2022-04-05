package main

import "fmt"

// declare a type constraint for Number interface
// want to constrain a type parameter to either int64 or float64, you can use this Number type constraint instead of writing out int64 | float64
type Number interface {
	int64 | float64 // union of int64 and float64 inside the interface
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
// In short Number instead of int64 | float64
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
// https://go.dev/doc/tutorial/generics
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s

	/*
		Declare a SumIntsOrFloats function with two type parameters (inside the square brackets), K and V, and one argument that uses the type parameters, m of type map[K]V. The function returns a value of type V.

		Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these, the comparable constraint is predeclared in Go. It allows any type whose values may be used as an operand of the comparison operators == and !=. Go requires that map keys be comparable. So declaring K as comparable is necessary so you can use K as the key in the map variable. It also ensures that calling code uses an allowable type for map keys.

		Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using | specifies a union of the two types, meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the calling code.

		Specify that the m argument is of type map[K]V, where K and V are the types already specified for the type parameters. Note that we know map[K]V is a valid map type because K is a comparable type. If we hadn’t declared K comparable, the compiler would reject the reference to map[K]V. */
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	/*
			Call the generic function you just declared, passing each of the maps you created.

		   Specify type arguments – the type names in square brackets – to be clear about the types that should replace type parameters in the function you’re calling.

		   As you’ll see in the next section, you can often omit the type arguments in the function call. Go can often infer them from your code.

		   Print the sums returned by the function. */
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// Call the generic function, omitting the type arguments
	// Go compiler can infer the types but not always possible
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	// Call function that using Number interface that include union int64 | float64
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
