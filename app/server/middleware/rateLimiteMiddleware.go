package middleware

import (
	"crac55/common/clog"
	"crac55/common/limiter"
	"net/http"
)

// LimitMiddleware 限流
func LimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetIPRateLimiter().GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			clog.Log().Infoln("Limit RemoteAddr:", r.RemoteAddr)
			return
		}
		next.ServeHTTP(w, r)
	})
}
