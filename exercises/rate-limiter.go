// A simple rate limiter, uses a queue (slice) to keep track of when the calls were made.
//
// Author: Ignacio Krichevsky

package exercises

import "time"

type RateLimiter struct {
	MaxCalls          int
	TimeWindowSeconds int
	callsQueue        []time.Time
}

func New(maxCalls int, timeWindowSeconds int) *RateLimiter {
	return &RateLimiter{
		maxCalls,
		timeWindowSeconds,
		make([]time.Time, maxCalls),
	}
}

func (rl *RateLimiter) IsWithinRate() bool {

	currentTime := time.Now()
	rl.callsQueue = append(rl.callsQueue, currentTime)

	if len(rl.callsQueue) <= rl.MaxCalls {
		return true
	}

	previousTime := rl.callsQueue[0]
	rl.callsQueue = rl.callsQueue[1:]

	return currentTime.Sub(previousTime).Seconds() > float64(rl.TimeWindowSeconds)

}
