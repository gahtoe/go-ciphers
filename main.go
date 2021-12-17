package main

import (
	"fmt"
	"github.com/gahtoe/go-ciphers/caesarcipher" 
	"github.com/gahtpe/go-ciphers/vigenerecipher"
)

func main() {
	const msg = "Text to encrypt."
	const shift = 3
	const key = "KEY"
	fmt.Println("CAESAR")
	encrypted := caesarcipher.Encrypt(msg, shift)
	decrypted := caesarcipher.Decrypt(encrypted, shift)
	fmt.Printf("Plaintext: %s\nShift: %d\nEncrypted: %s\nDecrypted: %s\n", msg, shift, encrypted, decrypted)
	fmt.Println("VIGENERE")
	encrypted := vigenerecipher.Encrypt(msg, key)
	decrypted := vigenerecipher.Decrypt(encrypted, key)
	fmt.Printf("Plaintext: %s\nKey: %s\nEncrypted: %s\nDecrypted: %s\n", msg, key, encrypted, decrypted)
}
