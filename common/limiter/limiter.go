package limiter

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var Limit = NewIPRateLimiter(15, 3)
var mu sync.Mutex

type IPRateLimiter struct {
	ips map[string]*visitor
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*visitor),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip].limiter = limiter
	i.ips[ip].lastSeen = time.Now()
	return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]
	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}
	limiter.lastSeen = time.Now()
	i.mu.Unlock()
	return limiter.limiter
}

func CleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range Limit.ips {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(Limit.ips, ip)
			}
		}
		mu.Unlock()
	}
}
