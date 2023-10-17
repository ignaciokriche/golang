// Author: Ignacio Krichevsky

package exercises

import (
	"math"
	"testing"
	"time"
)

func TestRateLimiterChannel(t *testing.T) {

	theTested := NewRateLimiterChannel(200 * time.Millisecond)
	defer theTested.Destroy()

	if !theTested.IsWithinRate() {
		t.Fatalf("1st call was expected to be in range.")
	}
	if theTested.IsWithinRate() {
		t.Fatalf("2nd call was not expected to be in range.")
	}

	time.Sleep(100 * time.Millisecond)
	if theTested.IsWithinRate() {
		t.Fatalf("3rd call was not expected to be in range.")
	}

	time.Sleep(110 * time.Millisecond)
	if !theTested.IsWithinRate() {
		t.Fatalf("4th call was expected to be in range.")
	}
	if theTested.IsWithinRate() {
		t.Fatalf("5th call was expected to be in range.")
	}

}

func TestRateIsLimited(t *testing.T) {

	theTested := NewRateLimiterChannel(100 * time.Millisecond)
	defer theTested.Destroy()

	inRangeCounter := 0
	stopCh := make(chan bool)

	go func(stopCh <-chan bool) {
		for {
			select {
			case <-stopCh:
				t.Log("go-routine stopping.\n")
				return
			default:
				if theTested.IsWithinRate() {
					inRangeCounter++
				}
			}

		}
	}(stopCh)

	t.Log(">>> sleeping...\n")
	time.Sleep(5 * time.Second)
	t.Log(">>> done sleeping.\n")
	stopCh <- true

	expected := 51
	if math.Abs(float64(inRangeCounter-expected)) > 3 {
		t.Fatalf("expecting: %v, was: %v.", expected, inRangeCounter)
	}

}
