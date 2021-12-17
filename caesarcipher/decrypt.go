package caesarcipher

func Decrypt(text string, shift uint) string {
	return Encrypt(text, 26 - (shift % 26))
}
