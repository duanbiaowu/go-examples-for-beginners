# 概述
`数组` 是具有相同数据类型的一组长度固定的数据项序列。其中数据类型可以是整型、布尔型等基础数据类型，也可以是自定义数据类型。
`数组长度` 必须是一个常量表达式，并且必须是一个非负数。同时，`数组长度` 也是数组类型的一部分，
例如 `[3]int` 和 `[5]int` 是不同的数组类型。 

# 语法规则
在不确定数组元素的具体数量时，可以通过指定 `长度` 分配空间。
```shell
var 变量名称 [长度]数据类型

# 例子
var arr [5]int
```

在确定数组元素的具体数量以及值时，可以省略 `长度` 变量，这种方式称为 `数组常量`。
```shell
var 变量名称 [...]数据类型{值1, 值2, 值3...}

# 例子
var arr [...]int{v1, v2, v3...}
```

# 获取值/改变值
和其他编程语言一样，**数组的元素可以通过 `索引` 来获取或修改，`索引` 从 `0` 开始。**

```go
package main

func main() {
	var arr [3]int

	arr[0] = 100 // 为数组第 1 个元素赋值
	arr[1] = 200 // 为数组第 2 个元素赋值
	arr[2] = 300 // 为数组第 3 个元素赋值

	println(arr[0]) // 输出数组第 1 个元素
	println(arr[1]) // 输出数组第 2 个元素
	println(arr[2]) // 输出数组第 3 个元素

	var arr2 = [...]int{400, 500, 600} // 使用 数组常量 定义
	println(arr2[0])                   // 输出数组第 1 个元素
	println(arr2[1])                   // 输出数组第 2 个元素
	println(arr2[2])                   // 输出数组第 3 个元素
}
// $ go run main.go
// 输出如下 
/**
    100
    200
    300
    400
    500
    600
 */
```

# 获取数组长度
通过 `len()` 函数获取。
```go
package main

import "fmt"

func main() {
	var arr [3]int

	arr[0] = 100 // 为数组第 1 个元素赋值
	arr[1] = 200 // 为数组第 2 个元素赋值
	arr[2] = 300 // 为数组第 3 个元素赋值

	fmt.Printf("数组长度 = %d\n", len(arr))
}
// $ go run main.go
// 输出如下 
/**
    数组长度 = 3
*/
```

# 遍历数组
可以使用两种方法遍历数组，[普通循环](for.md) 和 [range 循环](range.md)。

## 普通循环
```go
package main

import "fmt"

func main() {
	var arr [3]int

	arr[0] = 100 // 为数组第 1 个元素赋值
	arr[1] = 200 // 为数组第 2 个元素赋值
	arr[2] = 300 // 为数组第 3 个元素赋值

	for i := 0; i < len(arr); i++ {
		fmt.Printf("index = %d, val = %d\n", i, arr[i])
	}
}
// $ go run main.go
// 输出如下 
/**
    index = 0, val = 100
    index = 1, val = 200
    index = 2, val = 300
*/
```

## range 循环
```go
package main

import "fmt"

func main() {
	var arr [3]int

	arr[0] = 100 // 为数组第 1 个元素赋值
	arr[1] = 200 // 为数组第 2 个元素赋值
	arr[2] = 300 // 为数组第 3 个元素赋值

	for i, v := range arr {
		fmt.Printf("index = %d, val = %d\n", i, v)
	}
}
// $ go run main.go
// 输出如下 
/**
    index = 0, val = 100
    index = 1, val = 200
    index = 2, val = 300
*/
```

# 多维数组
将多个一维数组进行组装，得到一个多维数组。
```go
package main

import "fmt"

func main() {
	var arr [3][3]int // 定义一个 3 行 3 列的二维数组

	// 数组元素初始化
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = i*3 + j + 1
		}
	}

	// 初始化完成后，数组元素如下
	// -------------------
	// |  1  |  2  |  3  |
	// -------------------
	// |  4  |  5  |  6  |
	// -------------------
	// |  7  |  8  |  8  |
	// -------------------

	// 输出数组元素
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("第 %d 行，第 %d 列 的值 = %d\n", i+1, j+1, arr[i][j])
		}
	}
}
// $ go run main.go
// 输出如下 
/**
    第 1 行，第 1 列 的值 = 1
    第 1 行，第 2 列 的值 = 2
    第 1 行，第 3 列 的值 = 3
    第 2 行，第 1 列 的值 = 4
    第 2 行，第 2 列 的值 = 5
    第 2 行，第 3 列 的值 = 6
    第 3 行，第 1 列 的值 = 7
    第 3 行，第 2 列 的值 = 8
    第 3 行，第 3 列 的值 = 9
*/
```

# 扩展阅读
1. https://zh.wikipedia.org/wiki/%E6%95%B0%E7%BB%84
2. https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.1.md