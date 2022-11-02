# 概述

# 语法规则
调用 `rand 包` 即可，重要的一点是每次生成随机数之前，都设置随机数生成种子，否则可能每次生成的随机数都一样。

# 例子

## 随机生成数字
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 以当前时间的纳秒单位为种子

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int())
	}
}
// $ go run main.go
// 输出如下, 你的输出应该和这里的不一样
/**
    6322308781580164811
    8102638055079193560
    8689011158917073467
    6408490946268327546
    2346011052422006168
*/
```

## 随机生成指定区间数字
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano()) // 以当前时间的纳秒单位为种子
	r := rand.New(s)

	for i := 0; i < 5; i++ {
		fmt.Println(r.Intn(10))
	}
}
// $ go run main.go
// 输出如下, 你的输出应该和这里的不一样
/**
    5
    9
    7
    1
    3
*/
```