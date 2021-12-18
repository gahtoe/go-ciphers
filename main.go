package main

import (
	"fmt"
	c "github.com/gahtoe/go-ciphers/caesarcipher" 
	v "github.com/gahtoe/go-ciphers/vigenerecipher"
)

func main() {
	const msg = "Text to encrypt."
	const shift = 3
	const key = "KEY"
	fmt.Println("CAESAR")
	cencrypted := c.Encrypt(msg, shift)
	cdecrypted := c.Decrypt(cencrypted, shift)
	fmt.Printf("Plaintext: %s\nShift: %d\nEncrypted: %s\nDecrypted: %s\n\n", msg, shift, cencrypted, cdecrypted)
	fmt.Println("VIGENERE")
	vencrypted := v.Encrypt(msg, key)
	vdecrypted := v.Decrypt(vencrypted, key)
	fmt.Printf("Plaintext: %s\nKey: %s\nEncrypted: %s\nDecrypted: %s\n\n", msg, key, vencrypted, vdecrypted)
}
