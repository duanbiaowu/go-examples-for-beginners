# 普通打印
* 优点：内置函数，不需要引入额外的包，简单方便。
* 不足：无法进行格式化打印，无法完整打印复合数据结构 (如数组, Map 等)。

## println 函数
打印多个传入的参数，并自动加一个换行。

### 例子
```go
package main

func main() {
	println(1024, "hello world", true)
}
// $ go run main.go
// 输出如下 
/**
    1024 hello world true
*/
```

## print 函数
和 `println` 功能一样，但是不会自动加换行。

# 格式化打印
这里先介绍 2 个方法，分别是 `fmt` 包里面的 `Println()` 和 `Sprintf()`, 大多数场景下都适用。

## fmt.Println()
功能上和 `println 函数` 类似，但是可以打印复合数据结构 (如数组, Map 等)。

### 例子
```go
package main

import "fmt"

func main() {
	fmt.Println(1024, "hello world", true)
}

// $ go run main.go
// 输出如下 
/**
  1024 hello world true
*/
```

## fmt.Sprintf()
**最重要的格式化打印函数之一**，可以针对不同数据类型和数据结构进行打印，非常强大。

### 格式化规则

## 备注
笔者建议大家先记住 `fmt` 包里有这两个打印方法，具体的参数顺序和格式化规则可以暂时忽略， 
等后面学完了数组、Map 等复合数据结构以后，再根据具体场景回过头来查找对应的规则。
