package main

import (
	"fmt"

	"github.com/ctstewart/playfair/playfair"
)

func Example() {
	message := "lordgranvillesletter"
	key := "palmerston"
	encrypted := playfair.Encrypt(message, key)
	decrypted := playfair.Decrypt(encrypted, key)

	fmt.Println()
	fmt.Println("--- Example ---")
	fmt.Println()
	fmt.Printf("%12s %s\n", "Message:", message)
	fmt.Printf("%12s %s\n", "Key:", key)
	fmt.Printf("%12s %s\n", "Encrypted:", encrypted)
	fmt.Printf("%12s %s\n", "Decrypted:", decrypted)
}
