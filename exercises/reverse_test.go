package exercises

import "testing"

func TestRevertEmpty(t *testing.T) {
	if "" != Revert("") {
		t.Fail()
	}
}

func TestRevertOneChar(t *testing.T) {
	if "a" != Revert("a") {
		t.Fail()
	}
}

func TestRevertText(t *testing.T) {
	if "yksvehcirk oicangi" != Revert("ignacio krichevsky") {
		t.Fail()
	}
}

func TestRevertZero(t *testing.T) {
	if 0 != Revert(0) {
		t.Fail()
	}
}

func TestRevertNegative(t *testing.T) {
	number := -150
	expected := -51
	actual := Revert(number)
	if expected != actual {
		t.Fatalf("expecting %v. got: %v.", expected, actual)
	}
}

func TestRevertOne(t *testing.T) {
	if 1 != Revert(1) {
		t.Fail()
	}
}

func TestRevertTwentyOne(t *testing.T) {
	number := 21
	expected := 12
	actual := Revert(number)
	if expected != actual {
		t.Fatalf("expecting %v. got: %v.", expected, actual)
	}
}

func TestRevertNumber(t *testing.T) {
	number := 123456789
	expected := 987654321
	actual := Revert(number)
	if expected != actual {
		t.Fatalf("expecting %v. got: %v.", expected, actual)
	}
}
