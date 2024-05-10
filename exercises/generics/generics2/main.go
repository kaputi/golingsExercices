// generics2
// Make me compile!

package main

import "fmt"

type Number interface {
	int
}

func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(addNumbers(1.0, 2.3))
}

func addNumbers[T int | float64](n1, n2 T) T {
	return n1 + n2
}
