package trees

import (
	"errors"
	"slices"
	"testing"
)

func TestBfsLevelOrder(t *testing.T) {
	for _, test := range bfsLevelOrderTests {
		values, actualError := BfsLevelOrderTraversal(test.root)
		if !errors.Is(test.asserts.expectedError, actualError) {
			t.Errorf("test '%s' failed: with key '%d', expected error '%s', get error '%s'", test.asserts.name, test.asserts.key, test.asserts.expectedError, actualError)
		}
		if !slices.Equal(test.asserts.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, values)
		}
	}
}

func TestDfsInOrder(t *testing.T) {
	for _, test := range dfsInOrderTests {
		values := []int{}
		DfsInOrder(test.root, &values)
		if !slices.Equal(test.asserts.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, values)
		}
	}
}

func TestDfsPreOrder(t *testing.T) {
	for _, test := range dfsPreOrderTests {
		values := []int{}
		DfsPreOrder(test.root, &values)
		if !slices.Equal(test.asserts.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, values)
		}
	}
}

func TestDfsPostOrder(t *testing.T) {
	for _, test := range dfsPostOrderTests {
		values := []int{}
		DfsPostOrder(test.root, &values)
		if !slices.Equal(test.asserts.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.asserts.name, test.asserts.key, test.asserts.expected, values)
		}
	}
}
