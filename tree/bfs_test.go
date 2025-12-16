package tree

import (
	"errors"
	"slices"
	"testing"
)

func TestBfsLevelOrder(t *testing.T) {
	for _, test := range bfsLevelOrderTests {
		values, actualError := BfsLevelOrderTraversal(test.root)
		if !errors.Is(test.t.expectedError, actualError) {
			t.Errorf("test '%s' failed: with key '%d', expected error '%s', get error '%s'", test.t.name, test.t.key, test.t.expectedError, actualError)
		}
		if !slices.Equal(test.t.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.t.name, test.t.key, test.t.expected, values)
		}
	}
}
