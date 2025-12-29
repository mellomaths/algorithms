package search

import (
	"errors"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, test := range searchTests {
		actualValue, actualError := BinarySearch(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}

func TestBinarySearchFirstOccurrence(t *testing.T) {
	for _, test := range searchFirstOccurrenceTests {
		actualValue, actualError := BinarySearchFirstOccurrence(test.data, test.key)
		if actualValue != test.expected {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected '%d', get '%d'", test.name, test.data, test.key, test.expected, actualValue)
		}
		if !errors.Is(test.expectedError, actualError) {
			t.Errorf("test '%s' failed: input array '%v' with key '%d', expected error '%s', get error '%s'", test.name, test.data, test.key, test.expectedError, actualError)
		}
	}
}
