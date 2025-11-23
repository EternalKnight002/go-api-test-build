package greeter

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Doubles an integer input
func Double(n int) int {
	return n * 2
}