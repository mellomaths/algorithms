package strings

import "testing"

func TestIsNonTrivialRotation(t *testing.T) {
	for _, test := range nonTrivialRotationTests {
		actual := IsNonTrivialRotation(test.s1, test.s2)
		if actual != test.expected {
			t.Errorf("test '%s' failed: s1 '%s' s2 '%s', expected '%t', got '%t'", test.name, test.s1, test.s2, test.expected, actual)
		}
	}
}

func TestIsNonTrivialRotationIterative(t *testing.T) {
	for _, test := range nonTrivialRotationTests {
		actual := IsNonTrivialRotationIterative(test.s1, test.s2)
		if actual != test.expected {
			t.Errorf("test '%s' failed: s1 '%s' s2 '%s', expected '%t', got '%t'", test.name, test.s1, test.s2, test.expected, actual)
		}
	}
}
