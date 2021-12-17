package vigenerecipher

import (
	"strings"
)

func Sanitize(text string) string {
	var builder strings.Builder
	for _, r := range text {
		if 65 <= r && r <= 90 {
			builder.WriteRune(r)
		} else if 97 <= r && r <= 122 {
			builder.WriteRune(r - 32)
		}
	}	
	return builder.String() 
}
