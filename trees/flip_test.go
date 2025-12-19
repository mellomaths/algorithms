package trees

import (
	"slices"
	"testing"
)

func TestFlipTreeClockwise(t *testing.T) {
	for _, test := range flipTreeClockwiseTests {
		root, err := FlipTreeClockwise(test.root)
		if err != nil {
			t.Errorf("test '%s' failed: with key '%d', expected error '%s', got error '%s'", test.t.name, test.t.key, test.t.expectedError, err)
		}
		values, actualError := BfsLevelOrderTraversal(root)
		if actualError != nil {
			t.Errorf("test '%s' failed: with key '%d', BfsLevelOrderTraversal returned error '%s'", test.t.name, test.t.key, actualError)
		}
		if !slices.Equal(test.t.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', got '%d'", test.t.name, test.t.key, test.t.expected, values)
		}
	}
}
