package backoff

import (
	"time"
	"math/rand"
)

// BackoffJitter represents a backoff strategy with jitter.
type BackoffJitter func(baseBackoff time.Duration, retryCount int) time.Duration

// FullJitter applies full jitter to the backoff strategy.
func FullJitter(baseBackoff time.Duration, retryCount int) time.Duration {
	maxBackoff := baseBackoff * time.Duration(1<<retryCount) // Exponential backoff
	return time.Duration(rand.Int63n(int64(maxBackoff)))
}

// EqualJitter applies equal jitter to the backoff strategy.
func EqualJitter(baseBackoff time.Duration, retryCount int) time.Duration {
	maxBackoff := baseBackoff * time.Duration(1<<retryCount) // Exponential backoff
	halfMax := maxBackoff / 2
	return halfMax + time.Duration(rand.Int63n(int64(halfMax)))
}

// DecorrelatedJitter applies decorrelated jitter to the backoff strategy.
func DecorrelatedJitter(baseBackoff time.Duration, previousBackoff time.Duration) time.Duration {
	minBackoff := baseBackoff
	maxBackoff := previousBackoff * 3
	return time.Duration(rand.Int63n(int64(maxBackoff-minBackoff))) + minBackoff
}