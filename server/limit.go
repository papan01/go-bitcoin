package server

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

//訪客struct:
//limiter: 訪客訪問的限制器
//lastSeen: 訪客最後訪問頁面的最後時間
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

//使用map紀錄訪客的IP以及對應的struct
var visitors = make(map[string]*visitor)
var mu sync.Mutex

//背景執行cleanupVisitors
func init() {
	go cleanupVisitors()
}

//確認用戶是否存在於visitors中，若沒有則為其建立，更新最後訪問時間
//Limiter的核心概念使用Token_bucket-https://godoc.org/golang.org/x/time/rate#Limiter
//、https://en.wikipedia.org/wiki/Token_bucket
func getVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 10)
		visitors[ip] = &visitor{limiter, time.Now()}
		return limiter
	}
	v.lastSeen = time.Now()
	return v.limiter
}

//每分鐘確認是否有訪客超過三分鐘未在訪問，若有則從visitors中刪除
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		defer mu.Unlock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
	}
}

//提供使用限制器的主要接口
func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getVisitor(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
