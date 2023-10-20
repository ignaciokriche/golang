// Author: Ignacio Krichevsky

package exercises

import "testing"

func TestGreatestCommonDivisor(t *testing.T) {

	if GreatestCommonDivisor(0, 0) != 0 {
		t.Errorf("expecting 0.")
	}

	if GreatestCommonDivisor(8, 1) != 1 {
		t.Errorf("expecting 1.")
	}
	if GreatestCommonDivisor(1, 8) != 1 {
		t.Errorf("expecting 1.")
	}

	if GreatestCommonDivisor(12, 3) != 3 {
		t.Errorf("expecting 3.")
	}
	if GreatestCommonDivisor(3, 12) != 3 {
		t.Errorf("expecting 3.")
	}

	if GreatestCommonDivisor(21, 84) != 21 {
		t.Errorf("expecting 21.")
	}

	if GreatestCommonDivisor(82, 123) != 41 {
		t.Errorf("expecting 41.")
	}

	if GreatestCommonDivisor(54, 24) != 6 {
		t.Errorf("expecting 6.")
	}

}
