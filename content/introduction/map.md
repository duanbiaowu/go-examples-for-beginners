# 概述

`Map` 是一种键值对的无序集合，在其他编程语言中也被称为 `字典`, `Hash`, `关联数组`。

**重要的一点是: `Map 键` 的数据类型必须是可以比较的**，例如 `string`, `int`, `float64`，
[类型比较可以参考这篇文档](type_comparison.md) 。

`Map 值` 的数据类型可以是任意的。

# 语法规则

```shell
# 声明
var 变量名 map[键数据类型]值数据类型

# 声明及初始化
var 变量名 = make(map[键数据类型]值数据类型, 长度)   // 长度参数可以省略
```

# 获取值/改变值

```go
package main

import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Printf("Map 长度 = %d\n", len(m))

	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2

	fmt.Printf("Map 长度 = %d\n", len(m))

	fmt.Printf("zero = %T, %v\n", m["zero"], m["zero"])
	fmt.Printf("one = %T, %v\n", m["one"], m["one"])
	fmt.Printf("two = %T, %v\n", m["two"], m["two"])
}

// $ go run main.go
// 输出如下 
/**
  Map 长度 = 0
  Map 长度 = 3
  zero = int, 0
  one = int, 1
  two = int, 2
*/
```

# 删除元素

调用 `delete()` 方法完成。

```go
package main

import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Printf("Map 长度 = %d\n", len(m))

	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2

	fmt.Printf("Map 长度 = %d\n", len(m))

	delete(m, "one")
	delete(m, "two")

	fmt.Printf("Map 长度 = %d\n", len(m))
}

// $ go run main.go
// 输出如下 
/**
  Map 长度 = 0
  Map 长度 = 3
  Map 长度 = 1
*/
```

# 判断元素是否存在

```go
package main

func main() {
	var m = make(map[string]int)

	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2

	if _, ok := m["zero"]; ok {
		println(`m["zero"] 元素存在`)
	}

	delete(m, "zero")

	if _, ok := m["zero"]; !ok {
		println(`m["zero"] 元素不存在`)
	}
}

// $ go run main.go
// 输出如下 
/**
  m["zero"] 元素存在
  m["zero"] 元素不存在
*/
```

# 遍历 Map

**重要的一点是: Map 遍历是无序的。** 所以不能依赖于遍历的顺序，不论是 `键` 还是 `值`，
如果需要遍历时永远保持相同的顺序，需要提前将 `键` 做排序处理，参考 [有序 Map](sorted_map.md) 小节。

```go
package main

import "fmt"

func main() {
	var m = make(map[string]int)

	m["zero"] = 0
	m["one"] = 1
	m["two"] = 2

	for k, v := range m {
		fmt.Printf("key = %s, val = %d\n", k, v)
	}

	println("\n遍历 3 次，每次输出的结果可能不一样\n")
	for i := 0; i < 3; i++ {
		for k, v := range m {
			fmt.Printf("key = %s, val = %d\n", k, v)
		}
		fmt.Printf("第 %d 次遍历完成\n\n", i+1)
	}
}

// $ go run main.go
// 输出如下 
/**
  key = zero, val = 0
  key = one, val = 1
  key = two, val = 2

  遍历 3 次，每次输出的结果可能不一样

  key = one, val = 1
  key = two, val = 2
  key = zero, val = 0
  第 1 次遍历完成

  key = zero, val = 0
  key = one, val = 1
  key = two, val = 2
  第 2 次遍历完成

  key = one, val = 1
  key = two, val = 2
  key = zero, val = 0
  第 3 次遍历完成
*/
```

**备注:** 你的输出应该和这里不一样，多运行几次，看看是否每次都不一样。

# 并发不安全

最后要说明的很重要的一点是: **`Map` 不是并发安全的，** 也就是说，如果在多个线程中，同时对一个 Map 进行读写，会报错。
[互斥锁](mutex.md) 提供了一个简单的解决方案，后面会专门写一篇文档来说明如何才能 `并发安全`。

# Map 的 Key 不能是哪些类型？

- 映射（map）
- 切片
- 函数
- 包含不可比较字段的结构体类型
- 元素类型为不可比较类型的数组类型

在键类型的值之间必须可以施加操作符 == 和 != , 换句话说，**键类型的值必须要支持判等操作**。

**数组、切片和结构体不能作为 key (含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以。**

**如果要用结构体作为 key 可以提供 `Key()` 和 `Hash()` 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。**

```go
var badMap2 = map[interface{}]int{
	"1":   1,
	[]int{2}: 2, // 这里会引发panic。
	3:    3,
}

// map[type]type 基本类型时，不存在元素不会引起Panic，但统一判断可以作为最佳实践
set := map[string]struct{}{}
if set["hello"] {

}
```
