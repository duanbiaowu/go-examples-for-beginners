package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GeneralPipeline(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	// 执行了类似 Unix/Linux 命令 echo $nums | odd | sq | sum
	out := sum(sq(odd(echo(numbers))))
	assert.Equal(t, 35, <-out)
}

func Test_ProxyPipeline(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	out := proxy(numbers, echo, odd, sq, sum)
	assert.Equal(t, 35, <-out)
}

func Test_FanInOutPipeline(t *testing.T) {
	nums := makeRange(1, 10)
	in := echo(nums)

	var cs [4]<-chan int
	for i := range cs {
		cs[i] = sum(prime(in))
	}

	out := sum(merge(cs[:]))
	assert.Equal(t, 17, <-out)
}
