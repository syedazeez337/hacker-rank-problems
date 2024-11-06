/* count the number of words
in a camelCase string
*/

package main

import (
	"fmt"
	"unicode"
)

func camelCase(s string) int {
	wordCount := 1

	for _, char := range s {
		if unicode.IsUpper(char) {
			wordCount++
		}
	}

	return wordCount
}

func main() {
	s := "saveChangesInTheEditor"
	fmt.Println(camelCase(s))
}
