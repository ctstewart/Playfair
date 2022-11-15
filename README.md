# Playfair Cipher

## Overview

This project implements the Playfair cipher:

> The Playfair cipher was the first cipher to encrypt pairs of letters in cryptologic history.[2][3] Wheatstone invented the cipher for secrecy in telegraphy, but it carries the name of his friend Lord Playfair, first Baron Playfair of St. Andrews, who promoted its use.[3][4][5] The first recorded description of the Playfair cipher was in a document signed by Wheatstone on 26 March 1854.
>
> It was initially rejected by the British Foreign Office when it was developed because of its perceived complexity...It was however later used for tactical purposes by British forces in the Second Boer War and in World War I and for the same purpose by the British and Australians during World War II. This was because Playfair is reasonably fast to use and requires no special equipment - just a pencil and some paper...
>
> During World War II, the Government of New Zealand used it for communication among New Zealand, the Chatham Islands, and the coastwatchers in the Pacific Islands. Coastwatchers established by Royal Australian Navy Intelligence also used this cipher.
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
