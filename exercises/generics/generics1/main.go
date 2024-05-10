// generics1
// Make me compile!

package main

import "fmt"

func main() {
	print("Hello, World!")
	print(42)
}

func print[V int | string](value V) {
	fmt.Println(value)
}
