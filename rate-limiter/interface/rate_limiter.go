package rate_limiter

type RateLimiter interface {
	Limit() bool
}
