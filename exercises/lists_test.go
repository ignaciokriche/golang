// Author: Ignacio Krichevsky

package exercises

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesNil(t *testing.T) {

	list := stringToList("")
	list.RemoveDuplicates()

	var expected *ListNode = nil
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func TestRemoveDuplicatesLeft(t *testing.T) {

	list := stringToList("aab")

	list.RemoveDuplicates()

	expected := stringToList("ab")
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func TestRemoveDuplicatesMiddle(t *testing.T) {

	list := stringToList("abbc")

	list.RemoveDuplicates()

	expected := stringToList("abc")
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func TestRemoveDuplicatesRight(t *testing.T) {

	list := stringToList("abcc")

	list.RemoveDuplicates()

	expected := stringToList("abc")
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func TestRemoveDuplicatesNoDuplicates(t *testing.T) {

	list := stringToList("abcde")

	list.RemoveDuplicates()

	expected := stringToList("abcde")
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func TestRemoveMultipleDuplicates(t *testing.T) {

	list := stringToList("aaababjjbbccded")

	list.RemoveDuplicates()

	expected := stringToList("abjcde")
	if !reflect.DeepEqual(expected, list) {
		t.Fatalf("expecting: %v, got: %v.", expected, list)
	}

}

func stringToList(str string) *ListNode {
	if len(str) == 0 {
		return nil
	}
	return &ListNode{str[0:1], stringToList(str[1:])}
}
