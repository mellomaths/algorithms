package trees

import (
	"slices"
	"testing"
)

func TestDfsInOrder(t *testing.T) {
	for _, test := range dfsInOrderTests {
		values := []int{}
		DfsInOrder(test.root, &values)
		if !slices.Equal(test.t.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.t.name, test.t.key, test.t.expected, values)
		}
	}
}

func TestDfsPreOrder(t *testing.T) {
	for _, test := range dfsPreOrderTests {
		values := []int{}
		DfsPreOrder(test.root, &values)
		if !slices.Equal(test.t.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.t.name, test.t.key, test.t.expected, values)
		}
	}
}

func TestDfsPostOrder(t *testing.T) {
	for _, test := range dfsPostOrderTests {
		values := []int{}
		DfsPostOrder(test.root, &values)
		if !slices.Equal(test.t.expected, values) {
			t.Errorf("test '%s' failed: with key '%d', expected '%d', get '%d'", test.t.name, test.t.key, test.t.expected, values)
		}
	}
}
