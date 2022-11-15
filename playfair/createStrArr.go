package playfair

import (
	"fmt"
	"strings"
)

func createStrArr(str string) ([]string, error) {
	str = strings.ToUpper(str)

	// alphabet holds a list of A-Z
	seen_letters := map[string]bool{}

	var strArr []string
	for _, c := range str {
		char := string(c)

		if char == "J" {
			char = "I"
		}

		if _, was_seen := seen_letters[char]; was_seen {
			return []string{}, fmt.Errorf("duplicate letter in key: %s", char)
		}

		strArr = append(strArr, char)
		seen_letters[char] = true
	}

	for i := int('A'); i <= int('Z'); i++ {
		this_char := string(rune(i))
		if this_char == "J" {
			continue
		}
		if _, was_seen := seen_letters[this_char]; !was_seen {
			strArr = append(strArr, this_char)
		}
	}
	return strArr, nil
}
