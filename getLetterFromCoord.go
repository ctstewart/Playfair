package playfair

func getLetterFromCoord(x int, y int, l []string) string {
	letter := l[y*5+x]
	return letter
}
