package strings

import "strings"

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

func RotateStringNTimes(s string, n int) string {
	var sb strings.Builder
	i := 0
	for _, r := range s {
		sb.WriteString(s[1:])
		sb.WriteRune(r)
		s = sb.String()
		sb.Reset()
		i += 1
		if i == n {
			break
		}
	}
	return s
}
