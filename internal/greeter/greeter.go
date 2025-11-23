package greeter

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Doubles an integer input
func Double(n int) int {
	return n * 2
}

// Subtracts the second integer from the first
func Subtract(n1, n2 int) int {
	return n1 - n2
}

// Generates a unique, fun pet name by reversing and capitalizing the adjective and capitalizing the animal
func GeneratePetName(adjective, animal string) string {
	// 1. Reverse the adjective
	runes := []rune(adjective)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedAdjective := string(runes)

	// 2. Capitalize the first letter of the reversed adjective
	if len(reversedAdjective) > 0 {
		firstRune := []rune(reversedAdjective)[0]
		// Use strings.ToUpper for the first letter for simple ASCII capitalization
		upperFirst := strings.ToUpper(string(firstRune))
		reversedAdjective = upperFirst + reversedAdjective[len(string(firstRune)):]
	}

	// 3. Capitalize the animal
	capitalizedAnimal := strings.Title(strings.ToLower(animal))

	// 4. Combine
	return reversedAdjective + capitalizedAnimal
}