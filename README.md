# Playfair Cipher

## Overview

This project implements the Playfair cipher:

> The Playfair cipher or Playfair square or Wheatstone–Playfair cipher is a manual symmetric encryption technique and was the first literal digram substitution cipher. The scheme was invented in 1854 by Charles Wheatstone, but bears the name of Lord Playfair for promoting its use.
>
> The technique encrypts pairs of letters (bigrams or digrams), instead of single letters as in the simple substitution cipher and rather more complex Vigenère cipher systems then in use. The Playfair is thus significantly harder to break since the frequency analysis used for simple substitution ciphers does not work with it. The frequency analysis of bigrams is possible, but considerably more difficult. With 600 possible bigrams rather than the 26 possible monograms (single symbols, usually letters in this context), a considerably larger cipher text is required in order to be useful.
>
> -   [Wikipedia](https://en.wikipedia.org/wiki/Playfair_cipher)

## Usage

### Important

`message` below refers to a string you would like to encrypt/decrypt. **Only letters A-z are allowed. No spaces, special characters, numbers, etc.**

`key` has similar constraints, but also requires **no repeating letters.**

There are two options for usage:

### Run the Example

In your terminal:

```console
git clone https://github.com/ctstewart/playfair
cd playfair
go mod tidy
go run ./cmd
```

You'll then be prompted to either run the example or try encrypting/decrypting your own message.

If you choose the latter, don't forget to follow the constraints listed above.

### Import

You can import the Playfair package into your own project and use its public methods. Example:

```go
package your_package

import "github.com/ctstewart/playfair/playfair"

func main() {
    // Some code
}
```

You then have access to two methods:

`Encrypt(message, key)`

`Decrypt(message, key)`

`message` is the string you would like to encrypt/decrypt. For encryption, **only letters A-z are allowed. No spaces, special characters, numbers, etc.**

`key` is a string used for encryption and decryption. **It must be A-z, no spaces, special characters, numbers, etc., and have no repeating letters.**
