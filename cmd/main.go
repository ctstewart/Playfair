package main

import "fmt"

func main() {
	var eori string

	fmt.Println("Would you like to see an (e)xample or (t)ry for yourself?")
	fmt.Scanln(&eori)

	if eori == "e" {
		Example()
	} else if eori == "t" {
		Input()
	}
}
