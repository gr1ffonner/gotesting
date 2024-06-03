package strings

import (
	"strings"
)

func ReplaceChar(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
