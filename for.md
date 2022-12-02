# 概述

Go 仅提供了关键字 `for` 来表示循环，并没有提供 `while` 和 `do-while` 语句，这一点和主流编程语言不同。

# 语法规则

```shell
for 初始表达式; 条件表达式; 迭代表达式 {
    // do something
}
```

**注意: 迭代表达式中，不支持 `++i`, `--i` 这种形式，详情见 [自增/自减](inc_and_dec.md)。**

# 例子

## 单个计数器

```go
package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

// $ go run main.go
// 输出如下 
/**
  0
  1
  2
  3
  4
*/
```

## 多个计数器

```go
package main

func main() {
	for i, j := 1, 5; i <= 5; i, j = i+1, j-1 {
		println("i = ", i, " j = ", j)
	}
}

// $ go run main.go
// 输出如下 
/**
  i =  1  j =  5
  i =  2  j =  4
  i =  3  j =  3
  i =  4  j =  2
  i =  5  j =  1
*/
```

## 模仿 while

```go
package main

func main() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}

// $ go run main.go
// 输出如下 
/**
  0
  1
  2
  3
  4
*/
```

## 模仿 do-while

```go
package main

func main() {
	i := 0
	for {
		println(i)
		i++
		if i >= 5 {
			break
		}
	}
}

// $ go run main.go
// 输出如下 
/**
  0
  1
  2
  3
  4
*/
```

## 无限循环

```go
package main

func main() {
	i := 0
	for {
		println(i)
		i++
		if i >= 5 {
			break // 删除这行代码，将会进入无限循环
		}
	}
}

// $ go run main.go
// 输出如下 
/**
  0
  1
  2
  3
  4
*/

// 如果删除 `break` 语句，程序进入无限循环后可以使用 `Ctrl + C` 退出。
```