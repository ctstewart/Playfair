package playfair

import "strings"

func getCoordFromLetter(s string, l []string) [2]int {
	s = strings.ToUpper(s)
	var index int
	for i, letter := range l {
		if letter == s {
			index = i
			break
		}
	}
	y := index / 5
	x := index - (5 * y)
	return [2]int{x, y}
}
