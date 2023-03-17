package structural

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func Test_OpDecorator(t *testing.T) {
	double := func(n int) int {
		return n << 1
	}
	d := OpDecorator(double)
	assert.Equal(t, 10, d(5))

	increment := func(n int) int {
		return n + 1
	}
	inc := Operator(increment)
	assert.Equal(t, 6, inc(5))
}

func Test_TimedSumDecorator(t *testing.T) {
	normalSum := func(start, end int64) int64 {
		if start > end {
			start, end = end, start
		}
		var sum int64 = 0
		for i := start; i <= end; i++ {
			sum += i
		}
		return sum
	}

	nSum := timedSumDecorator(normalSum)
	assert.Equal(t, int64(500000500000), nSum(1, 1000000))

	gaussSum := func(start, end int64) int64 {
		if start > end {
			start, end = end, start
		}
		return (end - start + 1) * (end + start) / 2
	}

	gSum := timedSumDecorator(gaussSum)
	assert.Equal(t, int64(500000500000), gSum(1, 1000000))
}

func Test_HttpHandleDecorator(t *testing.T) {
	hello := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received Request %s from %s\n", r.URL.Path, r.RemoteAddr)
		_, _ = fmt.Fprintf(w, "hello world"+r.URL.Path)
	}

	Handler(hello, WithServerHeader, WithBasicAuth, WithAuthCookie, WithDebugLog)
}
