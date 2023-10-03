package exercises

import (
	"testing"
)

func TestSumDigitsEmpty(t *testing.T) {
	if SumDigits("") != 0 {
		t.Fatal("expecting 0.")
	}
}

func TestSumDigitsOneDigit(t *testing.T) {
	if SumDigits("5") != 5 {
		t.Fatal("expecting 5.")
	}
}

func TestSumDigitsTwoDigits(t *testing.T) {
	if SumDigits("55") != 1 {
		t.Fatal("expecting 1.")
	}

}

func TestSumDigitsThreeDigits(t *testing.T) {
	if SumDigits("849") != 3 {
		t.Fatal("expecting 3.")
	}
}

func TestSumDigitsMultipleDigits(t *testing.T) {
	if SumDigits("98401457874151") != 1 {
		t.Fatal("expecting 1.")
	}
}

func TestSumDigitsIterativeEmpty(t *testing.T) {
	if SumDigitsIterative("") != 0 {
		t.Fatal("expecting 0.")
	}
}

func TestSumDigitsIterativeOneDigit(t *testing.T) {
	if SumDigitsIterative("95") != 5 {
		t.Fatal("expecting 5.")
	}
}

func TestSumDigitsIterativeTwoDigits(t *testing.T) {
	if SumDigitsIterative("55") != 1 {
		t.Fatal("expecting 1.")
	}

}

func TestSumDigitsIterativeThreeDigits(t *testing.T) {
	if SumDigitsIterative("849") != 3 {
		t.Fatal("expecting 3.")
	}
}

func TestSumDigitsIterativeMultipleDigits(t *testing.T) {
	if SumDigitsIterative("98401457874151") != 1 {
		t.Fatal("expecting 1.")
	}
}
