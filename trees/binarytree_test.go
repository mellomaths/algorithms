package trees

import (
	"errors"
	"testing"
)

func TestMaximumHeightOfTree(t *testing.T) {
	for _, test := range maximumHeightOfTreeTests {
		height, err := MaximumHeightOfTree(test.root)
		if !errors.Is(err, test.asserts.expectedError) {
			t.Errorf("test '%s' failed: with key '%d', expected error '%s', get error '%s'", test.asserts.name, test.asserts.key, test.asserts.expectedError, err)
		}
		if height != test.asserts.expected {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, height)
		}
	}
}
