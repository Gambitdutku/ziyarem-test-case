package application

import (
	"sync"
	"time"
)

type CircuitBreaker struct {
	mu           sync.Mutex
	failCount    map[string]int
	openUntil    map[string]time.Time
	threshold    int
	openDuration time.Duration
}

func NewCircuitBreaker(threshold int, openDuration time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failCount:    make(map[string]int),
		openUntil:    make(map[string]time.Time),
		threshold:    threshold,
		openDuration: openDuration,
	}
}

func (c *CircuitBreaker) Allow(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if until, ok := c.openUntil[key]; ok {
		if time.Now().Before(until) {
			return false // hala açık
		}
		delete(c.openUntil, key) // reset işlevi
		c.failCount[key] = 0
	}
	return true
}

func (c *CircuitBreaker) Success(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failCount[key] = 0
}

func (c *CircuitBreaker) Failure(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.failCount[key]++
	if c.failCount[key] >= c.threshold {
		c.openUntil[key] = time.Now().Add(c.openDuration)
	}
}
