package vigenerecipher

import (
	"strings"
)

func decodePair(a, b rune) rune {
	return ((((a - 'A') + (b - 'A')) + 26) % 26) + 'A'
}

func Decrypt(text, key string) string {
	var builder strings.Builder
	stext, skey := Sanitize(text), Sanitize(key)
	for i, r := range stext {
		builder.WriteRune(decodePair(r, rune(skey[i%len(skey)])))
	}
	return builder.String()
}
