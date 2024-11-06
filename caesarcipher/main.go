package main

import "fmt"

func caesarCipher(s string, k int) string {
	// k = k % 26 ensures that k is within the range of 0 to 25
	k = k % 26
	result := []rune(s)

	// fmt.Println(result)

	for i, char := range result {
		if char >= 'a' && char <= 'z' {
			result[i] = 'a' + (char - 'a' + rune(k)) % 26
		} else if char >= 'A' && char <= 'Z' {
			result[i] = 'A' + (char - 'A' + rune(k)) % 26
		}
		// other characters remains the same
	}

	return string(result)
}

func main() {
	s := "#Hello-world$"
	k := 3
	res := caesarCipher(s, k)
	fmt.Println(res)
}
