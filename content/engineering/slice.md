---
title: 切片使用技巧
date: 2023-01-01
---

# 概述

Go 内置的 `append()` 和 `copy()` 两个函数非常强大，通过配合 `slice` 组合操作， 可以实现大多数 `容器类` 数据结构和基础算法，例如 `栈`, `队列` 的常规操作。

# 例子

## 追加元素

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = append(a, 11)

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9 10 11]
*/
```

## 追加 N 个元素

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = append(a, 11, 12, 13)

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9 10 11 12 13]
*/
```

## Copy 元素

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	b := make([]int, len(a))
	copy(b, a)

	fmt.Println(b)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9 10]
*/
```

## 删除一个元素

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	// 将元素 5 删除
	copy(a[4:], a[5:])
	a = a[:len(a)-1]

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 6 7 8 9 10]
*/
```

## 删除一段元素

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	// 将 [5, 6, 7] 3 个元素删除
	copy(a[4:], a[7:])
	a = a[:len(a)-3]

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 8 9 10]
*/
```

## 插入一个元素到指定位置

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	// 在 5 和 6 之间插入 1 个元素 0
	b := append(a, 0)
	copy(b[6:], b[5:])
	b[5] = 0

	fmt.Println(b)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 0 6 7 8 9 10]
*/
```

## 插入一段元素到指定位置

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	// 在 5 和 6 之间插入 3 个元素 [5, 6, 7]
	b := []int{5, 6, 7}
	a = append(a[:5], append(b, a[5:]...)...)

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 5 6 7 6 7 8 9 10]
*/
```

## 栈操作 Pop

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = a[:len(a)-1]

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9]
*/
```

## 栈操作 Push

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = append(a, 11)

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [1 2 3 4 5 6 7 8 9 10 11]
*/
```

## 队列操作 DeQueue

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = a[1:]

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [2 3 4 5 6 7 8 9 10]
*/
```

## 队列操作 EnQueue

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	a = append([]int{0}, a...)

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [0 1 2 3 4 5 6 7 8 9 10]
*/
```

## 元素反转

```go
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [10 9 8 7 6 5 4 3 2 1]
*/
```

## 洗牌算法

### 通过 rand.Intn() 实现

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [5 9 3 6 4 10 1 8 7 2]
*/
```

### 通过 rand.Shuffle() 实现

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}

	fmt.Println(a)

	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})

	fmt.Println(a)
}

// $ go run main.go
// 输出如下 
/**
  [1 2 3 4 5 6 7 8 9 10]
  [2 8 5 1 10 3 4 6 9 7]
*/
```

# Reference

1. https://github.com/golang/go/wiki/SliceTricks