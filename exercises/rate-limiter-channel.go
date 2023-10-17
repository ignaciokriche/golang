// A simple rate limiter, uses channel to limit the rate of calls.
// Based on: https://gobyexample.com/rate-limiting
//
// Author: Ignacio Krichevsky

package exercises

import "time"

type RateLimiterChannel struct {
	limiter   *time.Ticker
	firstCall bool
}

func NewRateLimiterChannel(interval time.Duration) *RateLimiterChannel {
	ticker := time.NewTicker(interval)
	return &RateLimiterChannel{limiter: ticker, firstCall: true}
}

func (rlc *RateLimiterChannel) Destroy() {
	rlc.limiter.Stop()
}

func (rlc *RateLimiterChannel) IsWithinRate() bool {
	defer func() { rlc.firstCall = false }()
	select {
	case <-rlc.limiter.C:
		return true
	default:
		return rlc.firstCall
	}

}
