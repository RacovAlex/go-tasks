package strings

import "strings"

func RepeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			var builder strings.Builder
			for j := 0; j < n/i; j++ {
				builder.WriteString(s[:i])
			}
			if builder.String() == s {
				return true
			}
		}
	}
	return false
}
