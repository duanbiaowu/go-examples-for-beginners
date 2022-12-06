# 概述

Go 中和时间相关的操作全部在 `time` 包。

# 语法规则

调用 `time` 包即可，重要的一点是: 不论将时间格式化为字符串，还是将字符串解析为时间，
**用到的时间参数固定为 `2006-01-02 15:04:05`** (至于为什么硬编码为这个时间，感兴趣的读者可以看看扩展阅读下面的文章),
而不是随意的时间参数，比如 `2018-08-08 18:28:38`。

# 例子

## 时间格式化

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006年01月02日 15时04分05秒"))
}

// $ go run main.go
// 输出如下, 你的输出应该和这里的不一样
/**
  2021-11-03 21:01:04
  2021年11月03日 21时01分04秒
*/
```

## 字符串解析为时间

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	s := `2021-11-03 21:01:04`
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		panic(err)
	}

	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006年01月02日 15时04分05秒"))
}

// $ go run main.go
// 输出如下
/**
  2021-11-03 21:01:04
  2021年11月03日 21时01分04秒
*/
```

## 获取年月日时分秒等属性

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Printf("%d\n", now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}

// $ go run main.go
// 输出如下, 你的输出应该和这里的不一样
/**
  2021
  11
  3
  21
  20
  7
*/
```

## 时间加减

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	afterHour := now.Add(time.Hour) // 当前时间加 1 个小时
	fmt.Println(afterHour.Format("2006-01-02 15:04:05"))

	beforeHour := now.Add(-time.Hour) // 当前时间减  1 个小时
	fmt.Println(beforeHour.Format("2006-01-02 15:04:05"))
}

// $ go run main.go
// 输出如下, 你的输出应该和这里的不一样
/**
  After a hour: 2021-11-03 21:36:17
  Before a hour: 2021-11-03 20:36:17

  After a year: 2022-11-03 21:40:43
  Before a year: 2020-11-03 21:40:43
*/
```

## 时间比较

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	afterHour := now.Add(time.Hour) // 当前时间加 1 个小时
	fmt.Printf("afterHour after now = %t\n", afterHour.After(now))

	beforeHour := now.Add(-time.Hour) // 当前时间减  1 个小时
	fmt.Printf("beforeHour before now = %t\n", beforeHour.After(now))
}

// $ go run main.go
// 输出如下
/**
  afterHour after now = true
  beforeHour before now = false
*/
```

# 扩展阅读

1. https://www.jianshu.com/p/c7f7fbb16932
2. https://jishuin.proginn.com/p/763bfbd62adb