package vigenerecipher

import (
	"strings"
)

func Sanitize(text string) string {
	var builder strings.Builder
	for _, r := range text {
		if 'A' <= r && r <= 'Z' {
			builder.WriteRune(r)
		} else if 'a' <= r && r <= 'z' {
			builder.WriteRune(r - 32)
		}
	}	
	return builder.String() 
}
