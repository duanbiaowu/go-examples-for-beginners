# 概述

`Map` 的遍历是无序的，这意味着不能依赖遍历的键值顺序。如果想实现 Map 遍历时顺序永远一致，
一个折中的方案时预先给 Map 的 `键` 排序，然后根据排序后的键序列遍历 Map, 这样可以保证每次遍历顺序都是一样的。

# 例子

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = make(map[int]string)

	m[0] = "zero"
	m[1] = "one"
	m[2] = "two"

	keys := make([]int, len(m)) // 将所有的键放入一个切片中
	index := 0
	for k, _ := range m {
		keys[index] = k
		index++
	}

	sort.Ints(keys) // 将所有的键进行排序

	for i := 0; i < 5; i++ {
		for _, key := range keys { // 根据排序后的键遍历 Map
			fmt.Printf("key = %d, val = %s\n", key, m[key])
		}
		fmt.Printf("第 %d 次遍历完成\n", i+1)
	}
}

// $ go run main.go
// 输出如下 
/**
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 1 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 2 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 3 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 4 次遍历完成
  key = 0, val = zero
  key = 1, val = one
  key = 2, val = two
  第 5 次遍历完成
*/
```

从输出的结果中可以看到，每次遍历的顺序都是一致的。