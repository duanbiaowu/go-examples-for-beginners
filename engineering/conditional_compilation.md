# 概述

`条件编译` 是指针对不同的平台，在编译期间选择性地编译特定的程序代码。
Go 通过引入 `build tag` 实现了条件编译。

# 例子

`条件编译` 一个常见的场景是: 针对同一个方法，在不同的环境中 (开发|测试|生产)，希望能输出不同等级的日志。

下面通过一个小例子来演示刚才描述的这种场景。

## go.mod

```shell
$ cat go.mod

# 输出如下
module helloworld

go 1.19
```

## foo 包

新建一个 `foo` 目录，并在目录下面建立 3 个文件: `debug.go`, `prod.go`, `main.go`。

### debug.go

将如下代码写入 `debug.go` 文件中。

```go
//go:build debug

package foo

func Mode() {
	println("Debug Mode")
}
```

### prod.go

将如下代码写入 `prod.go` 文件中。

```go
//go:build !debug

package foo

func Mode() {
	println("Production Mode")
}
```

### main.go

将如下代码写入 `main.go` 文件中。

```go
package main

import "helloworld/foo"

func main() {
	foo.Mode()
}
```

### 通过 tags 运行

`debug.go` 和 `prod.go` 两个文件中都有一个 `Mode` 方法，具体以哪个为准，需要编译时指定标签。通过指定不同的标签，运行程序可以得到不同的结果。

```shell
$ go run main.go
# 输出如下 
Production Mode

$ go run -tags debug main.go
# 输出如下
Debug Mode
```

# 多个编译条件

一个源文件中可以有多个 `build tags`，同一行的逗号隔开的 `tag` 之间是 `逻辑与` 的关系，空格隔开的 `tag` 之间是 `逻辑或` 的关系，
不同行之间的 `tag` 是 `逻辑与` 的关系。

### 示例

```shell
# 逻辑或，此源文件只能在 linux 或者 darwin 平台下编译
// +build linux darwin

# 逻辑与，此源文件只能在 linux/amd64 平台下编译
// +build amd64
// +build linux

# 此源文件只能在 linux/386 或者 darwin 平台下编译
// +build linux,386 darwin
```

# 小结

示例演示了通过 `条件编译` 运行代码，构建的方式也是一样的，只是将 `go run` 换成 `go build` 即可。
通过 `条件编译`，可以在构建服务时非常灵活地指定具体的环境、版本、部署方式等参数，真正做到 `一套代码，随地编译和运行`。