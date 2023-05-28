---
date: 2023-01-01
---

# defer

`defer` 语句经常用于成对的操作，比如 `打开文件/关闭文件` `连接网络/断开网络`, 合理地使用 `defer` 不仅可以提高代码可读性，也降低了忘记释放资源造成的泄漏等问题。

**正确使用 `defer` 语句的地方是在成功获取资之后**。

## 断开网络连接

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	
	// 此时资源有可能获取失败，执行 Close 导致 panic 
	// resp.Body.Close()
	
	if err != nil {
		panic(err)
	}
    
	defer func() {
		err = resp.Body.Close() // 关闭资源
		if err != nil {
			log.Fatal(err)
		}
	}()
}
```

## 关闭文件句柄

```go
package main

import (
	"os"
)

func main() {
	name := "/etc/hosts"
	file, err := os.Open(name)

	// 此时资源有可能获取失败，执行 Close 导致 panic
	// file.Close()
	
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close() // 关闭文件句柄
		if err != nil {
			panic(err)
		}
	}()

	hosts := make([]byte, 1024)
	_, err = file.Read(hosts)
	if err != nil {
		panic(err)
	}
}
```

## 计算程序耗时

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	
	// 错误的写法，等于注册 defer 函数的时候就已经计算好了输出值
	// defer fmt.Printf("executed time (%s)\n", time.Since(start))
	
	// 正确的写法
	defer func() {
		fmt.Printf("executed time (%s)\n", time.Since(start))
	}()

	time.Sleep(3 * time.Second) // 模拟程序耗时
}

// $ go run main.go
// 输出如下
/**
  executed time (3.0021534s)
*/
```

## defer 不会执行的情况

**`os.Exit` 会直接退出程序，不用调用已经注册的 `defer` 函数**。

```go
package main

import "os"

func main() {
	println("hello world")

	defer func() {
		println("hello defer")
	}()

	os.Exit(0)
}

// $ go run main.go
// 输出如下 
/**
  hello world
*/
```

通过输出结果可以看到，`defer` 并未执行，字符串 `hello defer` 没有输出，原因在于: 调用 `os.Exit()` 函数之后，程序会立即终止，
所有后面的代码和 `defer` 函数都不会执行。

# recover

`recover` 函数调用有着严格的要求，必须在 `defer` 函数中直接调用 `recover`，否则 `panic` 将无法被捕获。
如果 `defer` 函数中调用的是经过包装的 `recover` 函数，`panic` 将同样无法被捕获。

## recover 必须在 defer 中调用

### 错误的做法

```go
package main

import "fmt"

func main() {
	if r := recover(); r != nil { // 无法捕获到 panic
		fmt.Printf("panic = %v\n", r)
	}
	panic("some error")
}

// $ go run main.go
// 输出如下
/**
  panic: some error

  ...
  ...

  exit status 2

*/
```

### 正确的做法

将 `recover` 函数放置在 `defer` 函数中调用。

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil { // 可以捕获到 panic
			fmt.Printf("panic = %v\n", r)
		}
	}()

	panic("some error")
}

// $ go run main.go
// 输出如下 
/**
  panic = some error
*/
```

## recover 必须在 defer 中直接调用

### 错误的做法

```go
package main

import "fmt"

func myRecover() {
	if r := recover(); r != nil { // 无法捕获到 panic
		fmt.Printf("panic = %v\n", r)
	}
}

func main() {
	defer func() {
		myRecover()
	}()

	panic("some error")
}

// $ go run main.go
// 输出如下
/**
  panic: some error

  ...
  ...

  exit status 2

*/
```

**错误的原因在于:** `defer` 以匿名函数的方式运行，本身就等于包装了一层函数，
内部的 `myRecover` 函数包装了 `recover` 函数，等于又加了一层包装，变成了两
层包装，这时最外层的 `panic` 就无法被捕获了。

### 正确的做法 - 1

`defer` 直接调用 `myRecover` 函数，这样减去了一层包装，`panic` 就可以被捕获了。

```go
package main

import "fmt"

func myRecover() {
	if r := recover(); r != nil { // 无法捕获到 panic
		fmt.Printf("panic = %v\n", r)
	}
}

func main() {
	defer myRecover()

	panic("some error")
}

// $ go run main.go
// 输出如下 
/**
  panic = some error
*/
```

### 正确的做法 - 2

将 `recover` 函数放置在 `defer` 函数中调用，`panic` 就可以被捕获了。

```go
package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil { // 可以捕获到 panic
			fmt.Printf("panic = %v\n", r)
		}
	}()

	panic("some error")
}

// $ go run main.go
// 输出如下 
/**
  panic = some error
*/
```

## 多个 panic 只有一个被捕获

### 错误的做法

```go
package main

import "fmt"

func foo() {
	defer func() {
		println("recover 1")
		if err := recover(); err != nil { // 无法捕获到 panic
			fmt.Printf("[1] recovered %d\n", err)
		}
	}()

	defer func() {
		println("recover 2")
		if err := recover(); err != nil { // 无法捕获到 panic
			fmt.Printf("[2] recovered %d\n", err)
		}
	}()

	defer func() {
		println("recover 3")
		if err := recover(); err != nil { // 可以捕获到 panic
			fmt.Printf("[3] recovered %d\n", err)
		}
	}()

	defer func() {
		println("panic 1")
		panic(1)
	}()

	defer func() {
		println("panic 2")
		panic(2)
	}()

	defer func() {
		println("panic 3")
		panic(3)
	}()
}

func main() {
	foo()
}

// $ go run main.go
// 输出如下 
/**
  panic 3
  panic 2
  panic 1
  recover 3
  [3] recovered 1
  recover 2
  recover 1
*/
```

通过输出结果可以看到，即使抛出了多个 `panic`, 也只有最后一个被捕获。
因为第一个 `recover` 函数执行完后，会影响到后面的 `recover` 函数 (第一个 `recover` 捕获错误后，后面的 `recover` 不会捕获到任何错误)。

### 正确的做法

如果希望抛出的多个 `panic` 全部被捕获，应该在 `recover` 函数执行完后再依次执行 `panic`,
需要保证 `panic -> recover -> panic -> recover ...` 这样的链式关系。

```go
package main

import "fmt"

func foo() {
	defer func() {
		println("recover 1")
		if err := recover(); err != nil { // 可以捕获到 panic
			fmt.Printf("[1] recovered %d\n", err)
		}
	}()

	defer func() {
		println("recover 2")
		if err := recover(); err != nil { // 可以捕获到 panic
			fmt.Printf("[2] recovered %d\n", err)
			panic(err) // 捕获的同时继续抛出
		}
	}()

	defer func() {
		println("recover 3")
		if err := recover(); err != nil { // 可以捕获到 panic
			fmt.Printf("[3] recovered %d\n", err)
			panic(err) // 捕获的同时继续抛出
		}
	}()

	defer func() {
		println("panic 1")
	}()

	defer func() {
		println("panic 2")
	}()

	defer func() {
		println("panic 3")
		panic(3)
	}()
}

func main() {
	foo()
}

// $ go run main.go
// 输出如下 
/**
  panic 3
  panic 2
  panic 1
  recover 3
  [3] recovered 3
  recover 2
  [2] recovered 3
  recover 1
  [1] recovered 3
*/
```

# 小结

- 正确使用 `defer` 语句的地方是在成功获取资之后
- `os.Exit` 会直接退出程序，不用调用已经注册的 `defer` 函数
- `recover` 必须在 `defer` 函数中调用且必须直接调用
- 多个 `panic` 注册后，如果 `recover`, 那么只有 1 个 `panic` 会被捕获
