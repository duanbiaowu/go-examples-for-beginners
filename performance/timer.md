# 概述

`time.After` 和 `time.Tick` 不同，是一次性触发的，触发后 `timer` 本身会从时间堆中删除。
所以一般情况下直接用 `<-time.After` 是没有问题的， 不过在 for 循环的时候要注意:

# 每次分配新的 timer

```go
package performance

import (
	"testing"
	"time"
)

func Benchmark_Timer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-time.After(time.Millisecond):    // 每次生成新的 timer
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > slow.txt
```

# 复用一个 timer

刚才的示例代码中，每次进入 `select`，`time.After` 都会分配一个新的 `timer`。
因此会在短时间内创建大量的 `timer`，虽然 `timer` 在触发后会消失，但这种写法会造成无意义的 cpu 资源浪费。
正确的写法应该对 `timer` 进行复用。

```go
package performance

import (
	"testing"
	"time"
)

func Benchmark_Timer(b *testing.B) {
	timer := time.NewTimer(time.Second)

	for i := 0; i < b.N; i++ {
		timer.Reset(time.Millisecond) // 复用一个 timer
		select {
		case <-timer.C:
		}
	}
}
```

运行测试，并将基准测试结果写入文件:

```shell
# 运行 1000 次，统计内存分配
$ go test -run='^$' -bench=. -count=1 -benchtime=1000x -benchmem > fast.txt
```

# 使用 benchstat 比较差异

```shell
$ benchstat -alpha=100 fast.txt slow.txt 

# 输出如下:
name      old time/op    new time/op    delta
_TImer-8    1.27ms ± 0%    1.27ms ± 0%  +0.35%  (p=1.000 n=1+1)

name      old alloc/op   new alloc/op   delta
_TImer-8     0.00B        200.00B ± 0%   +Inf%  (p=1.000 n=1+1)

name      old allocs/op  new allocs/op  delta
_TImer-8      0.00           3.00 ± 0%   +Inf%  (p=1.000 n=1+1)
```

输出的结果分为了三行，分别对应基准期间的: 运行时间、内存分配总量、内存分配次数，采用了 `复用 timer` 方案后:
- 内存分配总量降至 0
- 内存分配次数降至 0

因为时间关系，基准测试只运行了 1000 次，运行次数越大，优化的效果越明显。感兴趣的读者可以将 `-benchtime` 调大后看看优化效果。

# 小结

通过复用 `time.After` 可以显著改善内存占用情况，在计时器比较多的业务场景中，还可以提升性能。