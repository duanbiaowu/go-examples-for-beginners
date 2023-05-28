---
date: 2023-01-01
---

# test-1

```go
package main

func foo(n int) (t int) {
	t = n
	defer func() {
		t += 3
	}()
	return t
}

func main() {
	println(foo(1))
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
4
````

### 结果分析

```go
package main

func foo(n int) (t int) {
	①t = n // 此时 t 为 1
	②defer func() {
    ④	t += 3 // 此时 t 为 4, 因为 t 是命名返回值，所以返回 4
	}()
	③return t // 此时 t 为 1
}

func main() {
	println(foo(1)) // 调用函数 foo(), 参数为 1
}
```

# test-2

```go
package main

func foo(n int) int {
	t := n
	defer func() {
		t += 3
	}()
	return t
}

func main() {
	println(foo(1))
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
1
````

### 结果分析

```go
package main

func foo(n int) int {
	①t := n // 此时 t 为 1
	②defer func() {
    ④	t += 3 // 此时 t 为 4, 但是 t 在作为返回值的时候等于 1, 所以这里的变化影响不到返回值
	}()
	③return t // 此时 t 为 1, 作为返回值返回，但是该函数不是命名返回值，所以这里直接返回 1
}

func main() {
	println(foo(1)) // 调用函数 foo(), 参数为 1
}
```

# test-3

```go
package main

func foo(n int) (t int) {
	defer func() {
		t += n
	}()
	return 1
}

func main() {
	println(foo(1))
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
2
````

### 结果分析

```go
package main

func foo(n int) (t int) {
	①defer func() {
    ③	t += n // 此时 t 为 2, 因为是命名返回值，所以改变了返回值，返回 2
	}()
	②return 1 // 此时 t 为 1
}

func main() {
	println(foo(1)) // 调用函数 foo(), 参数为 1
}
```

# test-4

```go
package main

func foo() (t int) {
	defer func(t int) {
		t += 5
	}(t)
	return 1
}

func main() {
	println(foo())
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
1
````

### 结果分析

```go
package main

func foo() (t int) {
	①defer func(t int) {	// 接收到的参数 t 为 0
    ④	t += 5	  // 此时 t 为 5, 但这里的 t 是 defer 函数的参数 t, 并不是 foo 函数的返回值的 t
    ②}(t)        // 此时 t 为 0, 参数为值传递
	③return 1    // 返回值为 1 
}

func main() {
	println(foo())
}

```

# test-5

```go
package main

func foo() (t int) {
	defer func(i int) {
		println(i)
		println(t)
	}(t)
	t = 1
	return 2
}

func main() {
	foo()
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
0
2
````

### 结果分析

```go
package main

func foo() (t int) {
	// 此时 t 为 0

	①defer func(n int) { // n 的值在注册时就已经决定了，等于传入的 t, 也就是 0
    ⑤	println(n) // 输出 0
    ⑥	println(t) // 此时已经执行完 return, 所以 t 为 2, 输出 2
    ②}(t)         // 注册 defer 函数时 t 为 0，传入的参数自然也是 0
	③t = 1    // 此时 t 为 1
	④return 2 // 此时 t 为 2, 因为 t 是函数命名返回值，return 执行后，开始执行 defer 函数
}

func main() {
	foo()
}
```

# test-6

```go
package main

func foo(index, value int) int {
	println(index)
	return index
}

func main() {
	defer foo(1, foo(3, 0))
	defer foo(2, foo(4, 0))
}
```

上面的代码会输出什么？思考之后 ...

```shell
$ go run main.go

# 输出如下
3
4
2
1
```

### 结果分析

```go
package main

func foo(index, value int) int {
	println(index)
	return index
}

func main() {
	defer ④foo(1, ①foo(3, 0))
	defer ③foo(2, ②foo(4, 0))
}
```

**4 个函数的先后执行顺序如下**:

1. 注册执行第 1 个 defer 函数: `foo(1, foo(3, 0))`
    - 第一个参数为 1, 第二个参数为调用 `foo(3, 0)` 的返回值
    - 计算第二个参数值
    - 调用 `foo(3, 0)`, 函数内部打印 3, 并且返回值为 3
    - 第二个参数值为 3
    - 调用 `defer foo(1, 3)` 完成注册，注意: 此时函数只是注册，但不会执行，所以不会打印第一个参数: 1

2. 注册执行第 2 个 defer 函数: `foo(2, foo(4, 0))`
    - 第一个参数为 2, 第二个参数为调用 `foo(4, 0)` 的返回值
    - 计算第二个参数值
    - 调用 `foo(4, 0)`, 函数内部打印 4, 并且返回值为 4
    - 第二个参数值为 4
    - 调用 `defer foo(2, 4)` 完成注册，注意: 此时函数只是注册，但不会执行，所以不会打印第一个参数: 2

3. 此时的 `defer 栈` 里面有两个函数
    - defer foo(1, 3)
    - defer foo(2, 4)

4. main 函数执行完成

5. 退出前开始执行 defer 函数 (先执行 `return`, 后执行 `defer`)
    - 先执行 `defer foo(2, 4)`， 函数内部打印 2
    - 再执行 `defer foo(1, 3)`， 函数内部打印 1


# 小结

本小结介绍了几种常见的 `defer` 函数求值问题，通过这些小例子，我们可以发现: 简单的 `defer` 函数经过编译器的包装后，处处是 "陷阱"，
这就要求我们要深入理解 `defer` 函数的运行机制，这样才不至于不经意间埋下了 `Bug`, 同时可以**通过一些工程代码约束来规避不必要的问题，
比如尽量避免在 `defer` 函数中定义复杂的参数和返回值，避免 `defer` 函数嵌套、避免多个 `defer` 函数之间毫无逻辑顺序等**。