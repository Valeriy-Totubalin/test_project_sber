package rate_limiter

import "time"

type TokenBucket struct {
	rate     int
	interval time.Duration
	time     time.Time
	tokens   int
}

func NewTokenBucket(rate int, interval time.Duration) RateLimiterInterface {
	return &TokenBucket{
		rate:     rate,
		interval: interval,
		time:     time.Now(),
		tokens:   rate,
	}
}

func (l *TokenBucket) CanDoWork() bool {
	currentTime := time.Now()
	// разница между временем последнего вызова и текущем
	dur := currentTime.Sub(l.time)
	// время за которое должен добавиться 1 токен
	onceRate := int(l.interval.Nanoseconds()) / l.rate
	// количество токенов, которые должны были появиться
	newTokens := int(dur.Nanoseconds()) / onceRate
	// добавляем эти токены
	if (l.tokens + newTokens) > l.rate {
		l.tokens = l.rate
	} else {
		l.tokens += newTokens
	}

	// если добавили токены, то обнулим время отчета
	if newTokens != 0 {
		l.time = currentTime
	}

	if l.tokens < 1 {
		return false
	}
	l.tokens -= 1

	return true
}
