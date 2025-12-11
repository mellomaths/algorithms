package strings

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for _, test := range palindromeTests {
		actual := IsPalindrome(test.input)
		if actual != test.expected {
			t.Errorf("test '%s' failed: input '%s', expected '%t', got '%t'", test.name, test.input, test.expected, actual)
		}
	}
}
