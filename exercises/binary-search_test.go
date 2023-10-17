// Author: Ignacio Krichevsky

package exercises

import "testing"

func TestBinarySearchEmpty(t *testing.T) {
	arr := []int{}
	expected := -1
	if index := BinarySearch(arr, 5); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchLen1NotFound(t *testing.T) {
	arr := []int{1}
	expected := -1
	if index := BinarySearch(arr, 2); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchLen1Found(t *testing.T) {
	arr := []int{1}
	expected := 0
	if index := BinarySearch(arr, 1); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchFound(t *testing.T) {

	arr := []int{1, 5, 6, 7, 8, 10, 15, 25}
	expected := 1
	if index := BinarySearch(arr, 5); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}

}

func TestBinarySearchNotFound(t *testing.T) {

	arr := []int{1, 5, 6, 7, 8, 10, 15, 25}
	expected := -1
	if index := BinarySearch(arr, 9); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}

}

func TestBinarySearchRecursiveEmpty(t *testing.T) {
	arr := []int{}
	expected := -1
	if index := BinarySearchRecursive(arr, 5); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchRecursiveLen1NotFound(t *testing.T) {
	arr := []int{1}
	expected := -1
	if index := BinarySearchRecursive(arr, 2); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchRecursiveLen1Found(t *testing.T) {
	arr := []int{1}
	expected := 0
	if index := BinarySearchRecursive(arr, 1); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}
}

func TestBinarySearchRecursiveFound(t *testing.T) {

	arr := []int{1, 5, 6, 7, 8, 10, 15, 25}
	expected := 1
	if index := BinarySearchRecursive(arr, 5); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}

}

func TestBinarySearchRecursiveNotFound(t *testing.T) {

	arr := []int{1, 5, 6, 7, 8, 10, 15, 25}
	expected := -1
	if index := BinarySearchRecursive(arr, 9); index != expected {
		t.Errorf("expecting: %v, was: %v.", expected, index)
	}

}
