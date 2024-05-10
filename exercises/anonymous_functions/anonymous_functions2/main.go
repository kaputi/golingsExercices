// anonymous functions2
// Make me compile!

package main

import "fmt"

func main() {

	func(n string) {
		fmt.Printf("Bye %s", n)
	}("Eduardo")
}
