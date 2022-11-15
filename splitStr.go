package playfair

func splitStr(s string) []string {
	var a []string
	for _, c := range s {
		char := string(c)
		a = append(a, char)
	}
	return a
}
