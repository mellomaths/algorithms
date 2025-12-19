package strings

import (
	"testing"
	"time"
)

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

func TestCompareIsNonTrivialRotationAlgorithms(t *testing.T) {
	for _, test := range nonTrivialRotationTests {
		start := time.Now() // Capture the start time
		actual := IsNonTrivialRotation(test.s1, test.s2)
		elapsed := time.Since(start) // Calculate the elapsed time
		t.Logf("IsNonTrivialRotation :: Processing time: %s\n", elapsed)
		start = time.Now() // Capture the start time
		actualI := IsNonTrivialRotationIterative(test.s1, test.s2)
		elapsed = time.Since(start) // Calculate the elapsed time
		t.Logf("IsNonTrivialRotationIterative :: Processing time: %s\n", elapsed)
		if actual != actualI {
			t.Errorf("test '%s' failed: s1 '%s' s2 '%s', IsNonTrivialRotation got '%t', IsNonTrivialRotationIterative got '%t'", test.name, test.s1, test.s2, actual, actualI)
		}
	}
}
