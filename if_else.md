# 概述

和其他编程语言中 `if/else` 规则一致，除了语法上略有差异。

# 语法规则

**`if` 和 `else if` 后面的条件表达式是不需要括号的。**

## 单个 if

```go
if condition {
	// do something	
}
```

### 例子

```go
package main

func main() {
	n := 1024
	if n > 0 {
		println("n > 0")
	}
}

// $ go run main.go
// 输出如下 
/**
  n > 0
*/
```

## 单个 if/else

```go
if condition {
	// do something	
} else {
	// do something	
}
```

### 例子

```go
package main

func main() {
	n := 1024
	if n > 0 {
		println("n > 0")
	} else {
		println("n <= 0")
	}
}

// $ go run main.go
// 输出如下 
/**
  n > 0
*/
```

## 多个分支

```go
if condition1 {
	// do something	
} else if condition2 {
	// do something else	
} else {
	// default
}
```

### 例子

```go
package main

func main() {
	n := 0
	if n > 0 {
		println("n > 0")
	} else if n < 0 {
		println("n < 0")
	} else {
		println("n = 0")
	}
}

// $ go run main.go
// 输出如下 
/**
  n == 0
*/
```