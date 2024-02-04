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

这里先介绍 2 个方法，分别是 `fmt` 包里面的 `Println()` 和 `Printf()`, 大多数场景下都适用。

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

## fmt.Printf()

**最重要的格式化打印函数之一**，可以针对不同数据类型和数据结构进行打印，非常强大。

## 格式化规则

**和 `C 系列` 编程语言的 `printf()` 格式化规则差不多。**

### 通用

* `%v`   默认格式
* `%+v`  针对结构体，在 `%v` 的基础上输出结构体的键名
* `%#v`  Go 语言语法格式的值
* `%T`   Go 语言语法格式的类型和值
* `%%`   输出 `%`, 相当于转义

### 整型

* `%b`	 二进制格式 
* `%c`	 对应的 Unicode 码 
* `%d`	 十进制 
* `%o`	 八进制 
* `%O`	 八进制，加上 `0o` 前缀 
* `%q`	 Go 语言语法转义后的单引号字符 (很少使用) 例如 97 会输出 `'a'` 
* `%x`	 十六进制 (小写), 例如 `0xaf` 
* `%X`	 十六进制 (大写), 例如 `0xAF` 
* `%U`	 Unicode 例如 `"U+%04X"`

### Bool

* `%t`   true 或 false

### 浮点型

* `%b`	 指数为 2 的幂的无小数科学计数法，例如 -123456p-78 
* `%e`	 科学计数法, 例如 -1.234456e+78 
* `%E`	 科学计数法, 例如 -1.234456E+78 
* `%f`	 常规小数点表示法 (一般使用这个), 例如 123.456 
* `%F`	 和 `%f` 功能一样

### 字符串

* `%s`	 字符串
* `%q`	 将双引号 `"` 转义后的字符串
* `%x`	 将字符串作为小写的十六进制
* `%X`	 将字符串作为大写的十六进制

### 指针

* `%p`	 地址的十六进制，前缀为 `0x`

### 例子

```go
package main

import "fmt"

func main() {
	n := 1024
	fmt.Printf("n = %d\n", n) // 输出整型

	pi := 3.1415
	fmt.Printf("pi = %f\n", pi) // 输出浮点数

	str := "hello world"
	fmt.Printf("str = %s\n", str) // 输出字符串

	yes := true
	fmt.Printf("yes = %t\n", yes) // 输出布尔型

	x := 17
	fmt.Printf("yes = %b\n", x) // 输出二进制
}

// $ go run main.go
// 输出如下
/**
    n = 1024
    pi = 3.141500
    str = hello world
    yes = true
    x = 10001
*/
```

## fmt.Printf() 技巧

在打印中，如果一个变量打印多次，可以通过 `[1]` 来表示后续变量全部以第一个为准。

### 例子

```go
package main

import (
	"fmt"
)

func main() {
	n := 1024
	fmt.Printf("%T %d %v\n", n, n, n)

	fmt.Printf("%T %[1]d %[1]v\n", n) // 可以使用 [1] 来表示引用第一个变量，这样只需要一个变量就可以了
}
// $ go run main.go
// 输出如下
/**
    int 1024 1024
    int 1024 1024
*/
```

## 备注

笔者建议大家先记住 `fmt` 包里有这两个打印方法，具体的参数顺序和格式化规则可以暂时忽略， 
等后面学完了数组、结构体、Map 等复合数据结构以后，再根据具体场景回过头来查找对应的规则。

# 附录

## 内置的`print`和`println`函数与`fmt`和`log`标准库包中相应的打印函数有什么区别？

1. 内置的 `print`/`println` 函数总是写入标准错误。 `fmt` 标准包里的打印函数总是写入标准输出。 `log` 标准包里的打印函数会默认写入标准错误，然而也可以通过 `log.SetOutput` 函数来配置。
2. 内置 `print`/`println` 函数的调用不能接受数组和结构体参数。
3. 对于组合类型的参数，内置的 `print`/`println` 函数将输出参数的底层值部的地址，而 `fmt` 和 `log` 标准库包中的打印函数将输出接口参数的动态值的字面形式。
4. 目前（Go 1.17），对于标准编译器，调用内置的 `print`/`println` 函数不会使调用参数引用的值逃逸到堆上，而 `fmt` 和 `log` 标准库包中的打印函数将使调用参数引用的值逃逸到堆上。
5. 如果一个实参有 `String() string` 或 `Error() string` 方法，那么 `fmt` 和 `log `标准库包里的打印函数在打印参数时会调用这两个方法，而内置的 `print`/`println` 函数则会忽略参数的这些方法。
6. 内置的 `print`/`println` 函数不保证在未来的 Go 版本中继续存在。
