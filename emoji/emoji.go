package emoji

import (
	"github.com/kaneshin/go-pkg/unicode"
)

// Contains ...
func Contains(s string) bool {
	for _, r := range []rune(s) {
		if unicode.IsEmoji(r) {
			return true
		}
	}
	return false
}
