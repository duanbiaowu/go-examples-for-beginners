package structural

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Operator func(int) int

type SumFunc func(int64, int64) int64

type HttpHandleDecorator func(http.HandlerFunc) http.HandlerFunc

func OpDecorator(fn Operator) Operator {
	return func(n int) int {
		result := fn(n)
		return result
	}
}

func getFuncName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}

func timedSumDecorator(f SumFunc) SumFunc {
	return func(start int64, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %d\n", getFuncName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

func WithBasicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithBasicAuth()")
		cookie, err := r.Cookie("Auth")
		if err != nil || cookie.Value != "Pass" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}

func WithDebugLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithDebugLog")
		_ = r.ParseForm()
		log.Println(r.Form)
		log.Println("path", r.URL.Path)
		log.Println("scheme", r.URL.Scheme)
		log.Println(r.Form["url_long"])
		for k, v := range r.Form {
			log.Println("key:", k)
			log.Println("val:", strings.Join(v, ""))
		}
		h(w, r)
	}
}

func Handler(h http.HandlerFunc, decors ...HttpHandleDecorator) http.HandlerFunc {
	for i := range decors {
		d := decors[len(decors)-i-1] // iterate in reverse
		h = d(h)
	}

	return h
}
