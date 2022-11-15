package main

import (
	"errors"
	"fmt"

	"github.com/ctstewart/playfair/playfair"
)

func Input() error {
	var eord string
	var message string
	var key string
	var ans string

	fmt.Println("Would you like to (e)ncrypt or (d)ecrypt?")
	fmt.Scanln(&eord)

	if eord != "e" && eord != "d" {
		return errors.New("please enter 'e' for encrypting or 'd' for decrypting")
	}

	fmt.Println("Please enter the key you would like to use:")
	fmt.Scanln(&key)

	if key == "" {
		return errors.New("please enter a valid key")
	}

	if eord == "e" {
		fmt.Println("Please enter the message you would like to encrypt:")
		fmt.Scanln(&message)
		ans = playfair.Encrypt(message, key)
	} else if eord == "d" {
		fmt.Println("Please enter the message you would like to decrypt:")
		fmt.Scanln(&message)
		ans = playfair.Decrypt(message, key)
	}

	fmt.Println(ans)

	return nil
}
