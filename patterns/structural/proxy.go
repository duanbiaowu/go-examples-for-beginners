package structural

import "net/http"

type server interface {
	handleRequest(string, string) (int, string)
}

type nginx struct {
	app             *application
	maxAllowRequest int
	rateLimiter     map[string]int
}

type application struct {
}

func NewNginxServer(maxReq int) *nginx {
	return &nginx{
		app:             &application{},
		maxAllowRequest: maxReq,
		rateLimiter:     make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	if allowed := n.checkRateLimiting(url); !allowed {
		return http.StatusForbidden, http.StatusText(http.StatusForbidden)
	}
	return n.app.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	// In the real environment, n.rateLimiter needs concurrency lock.
	if _, ok := n.rateLimiter[url]; !ok {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/list" && method == "GET" {
		return http.StatusOK, http.StatusText(http.StatusOK)
	}

	if url == "/app/create" && method == "POST" {
		return http.StatusCreated, http.StatusText(http.StatusCreated)
	}
	return http.StatusNotFound, http.StatusText(http.StatusNotFound)
}
