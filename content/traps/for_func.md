# 循环调用 defer 错误

`defer` 在函数退出时才会执行，在循环中执行 `defer` 释放资源时，由于延迟可能会引发 `资源泄露问题`。

## 错误的做法

```go
package main

import (
	"log"
	"os"
)

func main() {
	for i := 0; i < 5; i++ {
		f, err := os.Open("/path/to/file")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	zero := 0
	println(1 / zero) // 程序执行到这里异常退出，那么上面的循环中打开的 5 个文件句柄全部无法泄露
}
```

**错误的原因在于:** 极端情况下（比如 `for` 循环执行完程序异常，或者 `for` 还没执行完程序异常），将导致所有文件句柄无法释放，**造成资源泄露**。

再比如在第 4 次循环的时候，打开文件报错了，接着调用 `log.Fatal(err)` 结束程序，这时候，前面 3 次循环打开的 3 个文件句柄资源无法被释放，**造成资源泄露**。

## 正确的做法

**解决的方法:** 可以在 `for` 中构造一个局部函数，然后在局部函数内执行 `defer` 函数释放资源，
这样即使极端情况下程序异常退出，但是已经打开的文件句柄已经全部被释放，不会造成资源泄露。

```go
package main

import (
	"log"
	"os"
)

func main() {
	for i := 0; i < 5; i++ {
		func() {
			f, err := os.Open("/path/to/file")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
		}()
	}

	zero := 0
	println(1 / zero) // 程序执行到这里异常退出，但是上面的循环中打开的 5 个文件句柄资源已经全部释放，不会造成任何影响
}
```

# 循环调用 goroutine 错误

## 错误的做法

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		go func() {
			fmt.Printf("%v ", n)
		}()
	}

	time.Sleep(time.Second)
}

// $ go run main.go
// 输出如下
/**
  5 5 5 5 5
*/
```

**错误的原因在于**: 每次循环时 `goroutine` 的 `n` 都是同一个变量，循环结束后，该变量等于最后一次赋值后的值，也就是 `5` 。

## 正确的做法

下面描述了 2 种解决问题的方法。

**1. goroutine 函数参数**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		go func(val int) {
			fmt.Printf("%v ", val)
		}(n) // 将当前元素作为参数传入
	}

	time.Sleep(time.Second)
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样 
/**
  5, 4, 3, 1, 2
*/
```

**2. 变量重新赋值**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for _, n := range numbers {
		cur := n
		go func() {
			fmt.Printf("%v ", cur)
		}()
	}

	time.Sleep(time.Second)
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样 
/**
  5, 4, 1, 3, 2
*/
```