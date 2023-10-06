// Author: Ignacio Krichevsky

package exercises

import (
	"reflect"
	"testing"
)

func TestSumPairs(t *testing.T) {

	numbers := []int{1, 3, 4, 6, 3, 7, 10, 0, 9, -2}
	expectedPairs := []Pair{{3, 4}, {1, 6}, {7, 0}, {9, -2}}
	k := 7

	actualPairs := SumPairs(numbers, k)

	if !reflect.DeepEqual(expectedPairs, actualPairs) {
		t.Fatalf("expecting %v. got: %v.", expectedPairs, actualPairs)
	}

}

func TestDuplicatedSumPairs(t *testing.T) {

	numbers := []int{1, 3, 4, 3, 4, 6, 3, 7, 10, 0, 0, 7, 7}
	expectedPairs := []Pair{{3, 4}, {3, 4}, {1, 6}, {7, 0}, {0, 7}}
	k := 7

	actualPairs := SumPairs(numbers, k)

	if !reflect.DeepEqual(expectedPairs, actualPairs) {
		t.Fatalf("expecting %v. got: %v.", expectedPairs, actualPairs)
	}

}
