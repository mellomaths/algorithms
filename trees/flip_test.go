package trees

import (
	"errors"
	"slices"
	"testing"
)

func TestFlipTreeClockwise(t *testing.T) {
	for _, test := range flipTreeClockwiseTests {
		root, err := FlipTreeClockwise(test.root)
		if !errors.Is(err, test.asserts.expectedError) {
			t.Errorf("test '%s' failed: with key '%d', expected error '%s', got error '%s'", test.asserts.name, test.asserts.key, test.asserts.expectedError, err)
		}
		values, actualError := BfsLevelOrderTraversal(root)
		if !errors.Is(err, test.asserts.expectedError) {
			t.Errorf("test '%s' failed: with key '%d', BfsLevelOrderTraversal returned error '%s'", test.asserts.name, test.asserts.key, actualError)
		}
		if !slices.Equal(test.asserts.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', got '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, values)
		}
	}
}
