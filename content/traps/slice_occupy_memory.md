---
date: 2023-01-01
---

# 概述

切片的底层数据结构是数组，同样，切片的子切片会引用同样的数组。**如果切片不主动不释放的话，那么底层的数组就会一直占用着内存**。

## 切片返回值占用了整个数组

示例代码只是为了演示，没有任何实际意义。

### 错误的做法

```go
package main

import "fmt"

func getFirstThreeNumber() []byte {
	res := make([]byte, 1000)
	fmt.Println(len(res), cap(res))
	return res[:3]
}

func main() {
	res := getFirstThreeNumber()
	fmt.Println(len(res), cap(res))
}

// $ go run main.go
// 输出如下
/**
  1000 1000
  3 1000
*/
```

从输出结果中可以看到，即使函数已经返回切片，但是切片底层的数组一直被占用着，没有释放掉，浪费了很多内存。

### 正确的做法

分配一个合适大小的切片作为函数的返回值，这样函数返回后，切片底层的数组就会被释放掉。

```go
package main

import "fmt"

func getFirstThreeNumber() []byte {
	data := make([]byte, 1000)
	fmt.Println(len(data), cap(data))

	res := make([]byte, 3)
	copy(res, data[:3])
	return res
}

func main() {
	res := getFirstThreeNumber()
	fmt.Println(len(res), cap(res))
}

// $ go run main.go
// 输出如下
/**
  1000 1000
  3 3
*/
```

从输出结果中可以看到, 函数返回后，切片底层的数组已经被释放。