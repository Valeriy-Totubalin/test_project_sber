package rate_limiter

type RateLimiterInterface interface {
	CanDoWork() bool
}
