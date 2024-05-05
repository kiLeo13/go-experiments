package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

    slowServer := makeDelayedServer(20 * time.Millisecond);
    fastServer := makeDelayedServer(0)

    defer slowServer.Close()
    defer fastServer.Close()
    
    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    want := fastUrl
    got := Racer(slowUrl, fastUrl);

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}

func Racer(a, b string) (winner string) {
    aDuration := measureResponseTime(a)
    bDuration := measureResponseTime(b)

    if aDuration < bDuration {
        return a
    }
        
    return b
}

func measureResponseTime(url string) time.Duration {
    start := time.Now()
    http.Get(url)

    return time.Since(start)
}