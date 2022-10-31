# 自增
和主流编程语言的自增语法不同，Go 只支持 `i++` 方式，不支持 `++i` 方式。

**正确**
```go
package main

func main() {
	i := 1
	i++
	println(i)  // 输出 2
}
```

**错误**
```go
package main

func main() {
	i := 1
	++i     // 报错: '--' unexpected
	println(i)
}
```

# 自减
和主流编程语言的自减语法不同，Go 只支持 `i--` 方式，不支持 `--i` 方式。

**正确**
```go
package main

func main() {
	i := 1
	i--
	println(i)  // 输出 0
}
```

**错误**
```go
package main

func main() {
	i := 1
	--i     // 报错: '--' unexpected
	println(i)
}
```