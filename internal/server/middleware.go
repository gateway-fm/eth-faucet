package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v2"
	"github.com/kataras/hcaptcha"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni/v3"
)

type Limiter struct {
	mutex             sync.Mutex
	cache             *ttlcache.Cache
	proxyCount        int
	ttl               time.Duration
	ipWithdrawals     int
	walletWithdrawals int
}

func NewLimiter(proxyCount, ipWithdrawals, walletWithdrawals int, ttl time.Duration) *Limiter {
	if ipWithdrawals < 1 {
		ipWithdrawals = 1
	}
	if walletWithdrawals < 1 {
		walletWithdrawals = 1
	}
	cache := ttlcache.NewCache()
	cache.SkipTTLExtensionOnHit(true)
	return &Limiter{
		cache:             cache,
		proxyCount:        proxyCount,
		ttl:               ttl,
		ipWithdrawals:     ipWithdrawals,
		walletWithdrawals: walletWithdrawals,
	}
}

func (l *Limiter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	address, err := readAddress(r)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			renderJSON(w, claimResponse{Message: mr.message}, mr.status)
		} else {
			renderJSON(w, claimResponse{Message: http.StatusText(http.StatusInternalServerError)}, http.StatusInternalServerError)
		}
		return
	}

	if l.ttl <= 0 {
		next.ServeHTTP(w, r)
		return
	}

	clientIP := getClientIPFromRequest(l.proxyCount, r)
	l.mutex.Lock()
	if l.limitByKey(w, address, l.walletWithdrawals) || l.limitByKey(w, clientIP, l.ipWithdrawals) {
		l.mutex.Unlock()
		return
	}
	walletCount := l.increment(address)
	ipCount := l.increment(clientIP)
	l.mutex.Unlock()

	next.ServeHTTP(w, r)
	if w.(negroni.ResponseWriter).Status() != http.StatusOK {
		l.mutex.Lock()
		l.decrement(address)
		l.decrement(clientIP)
		l.mutex.Unlock()
		return
	}
	log.WithFields(log.Fields{
		"address":           address,
		"clientIP":          clientIP,
		"walletWithdrawals": fmt.Sprintf("%d/%d", walletCount, l.walletWithdrawals),
		"ipWithdrawals":     fmt.Sprintf("%d/%d", ipCount, l.ipWithdrawals),
	}).Info("Faucet withdrawal recorded")
}

func (l *Limiter) limitByKey(w http.ResponseWriter, key string, limit int) bool {
	if count, ttl, err := l.cache.GetWithTTL(key); err == nil && count.(int) >= limit {
		errMsg := fmt.Sprintf("You have exceeded the rate limit. Please wait %s before you try again", ttl.Round(time.Second))
		renderJSON(w, claimResponse{Message: errMsg}, http.StatusTooManyRequests)
		return true
	}
	return false
}

// increment bumps the withdrawal count for the key and returns the new count.
// When the key already exists it reuses the remaining TTL so the window stays
// anchored at the first claim (the cache is created with SkipTTLExtensionOnHit,
// so GetWithTTL returns the remaining time until expiry).
func (l *Limiter) increment(key string) int {
	if count, ttl, err := l.cache.GetWithTTL(key); err == nil {
		n := count.(int) + 1
		l.cache.SetWithTTL(key, n, ttl)
		return n
	}
	l.cache.SetWithTTL(key, 1, l.ttl)
	return 1
}

// decrement rolls back a withdrawal count, removing the entry once it reaches zero.
func (l *Limiter) decrement(key string) {
	if count, ttl, err := l.cache.GetWithTTL(key); err == nil {
		if c := count.(int); c > 1 {
			l.cache.SetWithTTL(key, c-1, ttl)
		} else {
			l.cache.Remove(key)
		}
	}
}

func getClientIPFromRequest(proxyCount int, r *http.Request) string {
	if proxyCount > 0 {
		xForwardedFor := r.Header.Get("X-Forwarded-For")
		if xForwardedFor != "" {
			xForwardedForParts := strings.Split(xForwardedFor, ",")
			// Avoid reading the user's forged request header by configuring the count of reverse proxies
			partIndex := len(xForwardedForParts) - proxyCount
			if partIndex < 0 {
				partIndex = 0
			}
			return strings.TrimSpace(xForwardedForParts[partIndex])
		}
	}

	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		remoteIP = r.RemoteAddr
	}
	return remoteIP
}

type Captcha struct {
	client *hcaptcha.Client
	secret string
}

func NewCaptcha(hcaptchaSiteKey, hcaptchaSecret string) *Captcha {
	client := hcaptcha.New(hcaptchaSecret)
	client.SiteKey = hcaptchaSiteKey
	return &Captcha{
		client: client,
		secret: hcaptchaSecret,
	}
}

func (c *Captcha) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if c.secret == "" {
		next.ServeHTTP(w, r)
		return
	}

	response := c.client.VerifyToken(r.Header.Get("h-captcha-response"))
	if !response.Success {
		renderJSON(w, claimResponse{Message: "Captcha verification failed, please try again"}, http.StatusTooManyRequests)
		return
	}

	next.ServeHTTP(w, r)
}
