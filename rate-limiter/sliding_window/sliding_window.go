package sliding_window

import "time"

type SlidingWindowRateLimiter struct {
	windowDuration time.Duration
	maxRequests    int
	requests       []time.Time
}

func NewSlidingWindowRateLimiter(windowDuration time.Duration, maxRequests int) SlidingWindowRateLimiter {
	return SlidingWindowRateLimiter{
		windowDuration: windowDuration,
		maxRequests:    maxRequests,
		requests:       make([]time.Time, 0),
	}
}

func (s *SlidingWindowRateLimiter) Limit() bool {
	now := time.Now()

	for len(s.requests) > 0 && now.Sub(s.requests[0]) > s.windowDuration {
		s.requests = s.requests[1:]
	}

	if len(s.requests) > s.maxRequests {
		return false
	}

	s.requests = append(s.requests, time.Now())

	return true
}
