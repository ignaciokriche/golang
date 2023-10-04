package exercises

import (
	"reflect"
	"testing"
)

func TestAnagramsEmpty(t *testing.T) {
	str := ""
	anagrams := Anagrams(str)
	expected := []string{}
	if !reflect.DeepEqual(expected, anagrams) {
		t.FailNow()
	}
}

func TestAnagramsOneChar(t *testing.T) {
	str := "a"
	anagrams := Anagrams(str)
	expected := []string{str}
	if !reflect.DeepEqual(expected, anagrams) {
		t.FailNow()
	}
}

func TestAnagramsTwoChars(t *testing.T) {
	str := "ab"
	anagrams := Anagrams(str)
	expected := []string{"ab", "ba"}
	if !reflect.DeepEqual(expected, anagrams) {
		t.FailNow()
	}
}

func TestAnagramsThreeChars(t *testing.T) {
	str := "abc"
	anagrams := Anagrams(str)
	expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	if !reflect.DeepEqual(expected, anagrams) {
		t.FailNow()
	}
}

func TestAnagrams(t *testing.T) {
	str := "zero"
	anagrams := Anagrams(str)
	expected := []string{"zero", "zeor", "zreo", "zroe", "zoer", "zore", "ezro", "ezor", "erzo", "eroz", "eozr",
		"eorz", "rzeo", "rzoe", "rezo", "reoz", "roze", "roez", "ozer", "ozre", "oezr", "oerz", "orze", "orez"}
	if !reflect.DeepEqual(expected, anagrams) {
		t.FailNow()
	}
}
