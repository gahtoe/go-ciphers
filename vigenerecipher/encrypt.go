package vigenerecipher

import (
	"strings"
)

func encodePair(a, b rune) rune {
	return ((((a - 'A') - (b - 'A')) % 26) + 26) + 'A'
}

func Encrypt(text, key string) string {
	var builder strings.Builder
	stext, skey := Sanitize(text), Sanitize(key)
	for i, r := range stext {
		builder.WriteRune(encodePair(r, rune(skey[i%len(skey)])))
	}
	return builder.String()
}
