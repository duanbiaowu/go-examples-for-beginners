---
title: 错误处理最佳实践
date: 2023-01-01
---

# 错误必须被处理

调用函数时，有很多函数总是成功返回，比如常见的 `println()` `len()`, 但是还有很多函数，因为各种不受控的影响 (比如 `网络中断`, `IO 错误` 等), 可能会调用失败甚至报错。
因此，**处理错误是程序中最重要的部分之一**。

Go 使用特定的类型 `error` 来标识错误，这和一些使用 `异常 (Exception)` 的编程语言不同。当调用函数发生错误时，一个约定俗成的做法是将 `错误值` 作为函数的最后一个返回值。
**如果函数返回一个错误时，调用方必须处理该错误，而不能想当然地认为函数执行成功，忽略错误**。

## 错误没有被处理导致的报错

### 错误的做法

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 模拟发起一个错误请求
	// 并且没有处理错误
	resp, _ := http.Get("localhost:3306") 
	defer resp.Body.Close()

	code := resp.StatusCode
	fmt.Printf("Http Code = %d\n", code)

	ct := resp.Header.Get("Content-Type")
	fmt.Printf("Content-Type = %s\n", ct)
}

// $ go run main.go
// 输出如下 
/**
  panic: runtime error: invalid memory address or nil pointer dereference
  ...
  ...
  exit status 2
*/
```

在上述示例代码中，向一个本地不存在的服务发起 HTTP 请求 (返回值肯定会报错)，
但是程序并没有处理错误，所以后续的读取操作和资源关闭操作就会报错。

### 正确的做法

如果函数返回一个错误时，调用方必须处理该错误，在获得返回值之后，要第一时间处理。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("localhost:3306")
	// 第一时间处理错误
	if err != nil {
		fmt.Printf("HTTP GET %s\n", err)
		return
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	code := resp.StatusCode
	fmt.Printf("Http Code = %d\n", code)

	ct := resp.Header.Get("Content-Type")
	fmt.Printf("Content-Type = %s\n", ct)
}

// $ go run main.go
// 输出如下，提前处理了错误并退出
/**
  HTTP GET Get "localhost:3306": unsupported protocol scheme "localhost"
*/
```

## 小结

笔者在学习和使用 `Go` 语言期间，能够深刻感受到 **语言本身对于错误处理的鼓励态度**，这不仅可以使调用方更快速地理解上下文，也可以帮助开发人员构建更加健壮、可维护性的代码。

# 优雅地处理错误

程序中过多的 `错误处理逻辑` 会让代码变得臃肿不堪，阅读时将很难分辨哪些是正常的程序逻辑，哪些是错误处理，一个好的办法是 **将错误处理的部分抽象封装起来**。

## 示例

通过以下几个常规的文件操作来做演示:

- 创建一个文件
- 写入一些字符串
- 读取一些字符串
- 关闭文件
- 删除文件

### 代码可读性较差

```go
package main

import (
	"log"
	"os"
)

func main() {
	name := "/tmp/error_handle.log"

	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}

		err = os.Remove(name)
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = file.WriteString("hello world")
	if err != nil {
		log.Fatal(err)
	}

	str := make([]byte, 1024)
	_, err = file.Read(str)
	if err != nil {
		log.Fatal(err)
	}
}
```

在上面的示例代码中，针对各种文件操作可能引起的错误，主体程序分别进行了处理，在这个短小的程序中，
有将近 `70%` 的代码是在处理异常，整体代码读上去比较混乱，无法捕捉到核心的处理逻辑在什么地方。

### 提高代码可读性

通过将错误处理部分封装为一个 `操作函数`，程序主体部分只需要关注 `操作函数` 的返回值即可，
不需要再对各种文件操作可能引起的错误分别处理，提高了代码的可读性，降低了代码复杂性和调用方的心智负担。

```go
package main

import (
	"log"
	"os"
)

// 将处理部分封装为一个函数
func fileBaseOperate(name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return
	}

	defer func() {
		err = file.Close()
		if err != nil {
			return
		}
		err = os.Remove(name)
	}()

	_, err = file.WriteString("hello world")
	if err != nil {
		return
	}

	str := make([]byte, 1024)
	_, err = file.Read(str)

	return
}

func main() {
	// 调用方只关注封装函数即可
	err := fileBaseOperate("/tmp/error_handle.log")
	if err != nil {
		log.Fatal(err)
	}
}
```

## 自定义错误类型

### 较差的实现方案

```go
package main

import (
	"errors"
	"log"
	"time"
)

type Transaction struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
}

// NewTransaction 创建新订单
func NewTransaction(id string) (*Transaction, error) {
	if len(id) == 0 {
		return nil, errors.New("transaction id can not be empty")
	}

	return &Transaction{
		ID:        id,
		Amount:    0,
		CreatedAt: time.Now(),
	}, nil
}

func main() {
	t, err := NewTransaction("")
	if err != nil && err.Error() == "transaction id can not be empty" {
		log.Fatal(err)
	} else {
		log.Printf("transaction id = %s", t.ID)
	}
}
```

```shell
$ go run main.go
# 输出如下 
2022/11/17 21:41:23 transaction id can not be empty
exit status 1
```

在上述示例代码中，调用方以 `硬编码` 的方式来处理错误，这种方式非常不利于扩展，可以通过 `自定义错误类型` 来改进。

### 更好的实现方案

`自定义错误类型` 比较好的命名实践是: **对于存储为全局变量的错误值，根据是否导出，使用前缀 `Err` 或  `err`。对于自定义错误类型，改用后缀 `Error`** 。

```go
package main

import (
	"errors"
	"log"
	"time"
)

var (
	// TransIDEmptyErr 自定义错误类型
	TransIDEmptyErr = errors.New("transaction id can not be empty")
)

type Transaction struct {
	ID        string
	Amount    float64
	CreatedAt time.Time
}

// NewTransaction 创建新订单
func NewTransaction(id string) (*Transaction, error) {
	if len(id) == 0 {
		return nil, TransIDEmptyErr
	}

	return &Transaction{
		ID:        id,
		Amount:    0,
		CreatedAt: time.Now(),
	}, nil
}

func main() {
	t, err := NewTransaction("")
	if err != nil && errors.Is(err, TransIDEmptyErr) {
		log.Fatal(err)
	} else {
		log.Printf("transaction id = %s", t.ID)
	}
}
```

```shell
$ go run main.go
# 输出如下 
2022/11/17 21:49:23 transaction id can not be empty
exit status 1
```

在上述示例代码中，通过 `自定义错误类型`，可以避免调用方以 `硬编码` 的方式来处理错误。

# 扩展阅读 

**如何区分 panic 和 error 两种使用方式 ?** 

一个约定俗成的方式是: 导致关键流程出现不可修复性错误时使用 `panic`, 其他情况使用 `error`。 
另外，包内部总是应该从 `panic` 中 `recover`, 不允许超出包范围的显式 `panic`, 包提供给调用者的 API 应该返回 `error`。