package vigenerecipher

import (
	"strings"
	
	"github.com/gahtoe/go-ciphers/util"
)

func encodePair(a, b rune) rune {
	return ((((a - 'A') - (b - 'A')) % 26) + 26) + 'A'
}

func Encrypt(text, key string) string {
	var builder strings.Builder
	stext, skey := util.Sanitize(text), util.Sanitize(key)
	for i, r := range stext {
		builder.WriteRune(encodePair(r, rune(skey[i%len(skey)])))
	}
	return builder.String()
}
