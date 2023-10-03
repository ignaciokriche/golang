// Author: Ignacio Krichevsky

package exercises

import (
	"testing"
	"time"
)

func TestGivenMaxTenCallsForOneSecondInterval_WhenCallingOnce_ThenMustBeWithinRate(t *testing.T) {
	var theTested *RateLimiter = New(10, 1) // 10 calls in 1 second.
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
}

func TestGivenMaxTenCallsForOneSecondInterval_WhenCallingTenTimes_ThenMustBeWithinRate(t *testing.T) {
	var theTested *RateLimiter = New(10, 1) // 10 calls in 1 second.
	for i := 0; i < 10; i++ {
		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}
	}
}

func TestGivenMaxTenCallsForOneSecondInterval_WhenCallingElevenTimes_ThenLastCallMustNotBeWithinRate(t *testing.T) {

	var theTested *RateLimiter = New(10, 1) // 10 calls in 1 second.

	for i := 0; i < 10; i++ {
		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}
	}

	if theTested.IsWithinRate() {
		t.Fatal("expecting false")
	}

}

func TestUptoFiveCallsWithinTwoSecondsMustBeInRange(t *testing.T) {

	var theTested *RateLimiter = New(5, 2) // 5 calls in 2 seconds.

	for tries := 2; tries > 0; tries-- {

		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}
		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}

		time.Sleep(1 * time.Second)

		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}

		time.Sleep(500 * time.Millisecond)

		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}

		if !theTested.IsWithinRate() {
			t.Fatal("expecting true")
		}

		time.Sleep(100 * time.Millisecond)

		if theTested.IsWithinRate() {
			t.Fatal("expecting false")
		}

		time.Sleep(2 * time.Second)
	}

}

func TestGivenFourCalls_WhenMakingFiveCallsTwoSecondsLater_ThenMustBeWithinRange(t *testing.T) {

	var theTested *RateLimiter = New(5, 2) // 5 calls in 2 second.

	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}

	time.Sleep(2 * time.Second)

	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}

	if theTested.IsWithinRate() {
		t.Fatal("expecting false")
	}

}

func TestLastThreeMustBeWithinRange(t *testing.T) {

	var theTested *RateLimiter = New(5, 1) // 5 calls in 1 second.

	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}

	time.Sleep(800 * time.Millisecond)

	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}

	time.Sleep(800 * time.Millisecond)

	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}
	if !theTested.IsWithinRate() {
		t.Fatal("expecting true")
	}

	if theTested.IsWithinRate() {
		t.Fatal("expecting false")
	}

}
