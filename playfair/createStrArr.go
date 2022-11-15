package playfair

import (
	"sort"
	"strings"
)

func createStrArr(str string) []string {
	str = strings.ToUpper(str)
	alphabet := map[string]int{
		"A": 65,
		"B": 66,
		"C": 67,
		"D": 68,
		"E": 69,
		"F": 70,
		"G": 71,
		"H": 72,
		"I": 73,
		"K": 75,
		"L": 76,
		"M": 77,
		"N": 78,
		"O": 79,
		"P": 80,
		"Q": 81,
		"R": 82,
		"S": 83,
		"T": 84,
		"U": 85,
		"V": 86,
		"W": 87,
		"X": 88,
		"Y": 89,
		"Z": 90,
	}
	var strArr []string
	for _, character := range str {
		char := string(character)
		if char == "J" {
			strArr = append(strArr, "I")
		} else {
			strArr = append(strArr, char)
		}
		delete(alphabet, char)
	}
	var alphaArr []int
	for _, character := range alphabet {
		alphaArr = append(alphaArr, character)
	}
	sort.Ints(alphaArr)
	for _, character := range alphaArr {
		char := string(rune(character))
		strArr = append(strArr, char)
	}
	return strArr
}
