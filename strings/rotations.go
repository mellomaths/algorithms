package strings

import "strings"

// IsNonTrivialRotation checks if s2 is a non-trivial rotation of s1.
// A non-trivial rotation means that s2 is not equal to s1 and
// can be obtained by rotating s1.
func IsNonTrivialRotation(s1 string, s2 string) bool {
	if s1 == s2 {
		return false
	}
	return strings.Contains((s1 + s1), s2)
}

func IsNonTrivialRotationIterative(s1 string, s2 string) bool {
	if s1 == s2 {
		return false
	}
	var sb strings.Builder
	for _, r := range s1 {
		sb.WriteString(s1[1:])
		sb.WriteRune(r)
		s1 = sb.String()
		sb.Reset()
		if s1 == s2 {
			return true
		}
	}
	return false
}
