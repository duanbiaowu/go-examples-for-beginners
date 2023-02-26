# 概述

> 工欲善其事，必先利其器。

优秀的工具配合熟练的使用，往往可以让开发效率大幅度提升，本小节介绍 Go 里面经常使用到的命令行工具。

# install

> `go install` 命令编译并安装指定的包以及对应的依赖包。

```shell
# 安装 golint 包
$ go install golang.org/x/lint/golint@latest
# go: downloading golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7

# 一般会直接将命令放入 $GOPATH/bin
$ which golint
# /home/codes/go/bin/golint
```

# get

> `go get` 命令将指定的包以及对应的依赖包加入到当前 `module`。

`go get` 和 `go install` 主要区别在于: `install` 是命令的全局安装，不会将包及其依赖加入到当前 `module`。

**需要注意的一点是: 每个包都有对应的 Go 版本以及其他包依赖，如果指定了包的版本号，但是当前 Go 版本或者依赖包的版本不满足条件，将无法安装**。

## 添加最新可用包

```shell
# 获取 golint 包, -u 参数表示获取指定的包的依赖项，以便在包有新的版本可用时使用
# 如果包名称后面不加 `@版本号`，则默认为 `latest` 最新可用的
$ go get -u golang.org/x/lint/golint
# go: added golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
# go: added golang.org/x/tools v0.3.0

# go.mod 中依赖项已经更新 
$ cat go.mod

# 输出如下:
module helloworld

go 1.19

require (
        golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
        golang.org/x/tools v0.3.0 // indirect
)
```

## 添加指定可用包

```shell
# 指定包的版本号为 @v1.4.1 
$ go get -u github.com/spf13/cast@v1.4.1 
# go: added github.com/spf13/cast v1.4.1

# go.mod 中依赖项已经更新 
$ cat go.mod

# 输出如下:
module helloworld

go 1.19

require (
	github.com/spf13/cast v1.4.1 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.3.0 // indirect
)
```

# mod

`go mod` 命令管理 `module` 相关操作。

## 子命令

```shell
go mod init  # 初始化
go mod tidy  # 更新 (移除) 依赖文件
go mod download  # 下载依赖文件

go mod vendor  # 将依赖全部归档到 vendor 目录
go mod edit    # 修改依赖文件
go mod graph   # 打印依赖关系 图
go mod verify  # 校验依赖
```

# clean

> `go clean` 命令删除执行其他命令时产生的目标和缓存文件。

# list

> `go list` 命令展示 `module` 或包的信息。

## 示例

```shell
# 不带参数，展示 module
$ go list
# helloworld

# 展示某个包的信息
$ go list golang.org/x/lint/golint
# 默认只输出包名
# golang.org/x/lint/golint

# 获取详细信息，并以 json 格式输出
$ go list -json golang.org/x/lint/golint
{
    "Dir": "/home/codes/go/pkg/mod/golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616/golint",
    "ImportPath": "golang.org/x/lint/golint",
    "Name": "main",
  ...
  ...
}

# 获取包的路径
$ go list -f '{{.Dir}}' golang.org/x/lint/golint
# /home/codes/go/pkg/mod/golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616/golint

# 切换路径到包的目录
$ cd $(go list -f '{{.Dir}}' golang.org/x/lint/golint)

$ pwd
# /home/codes/go/pkg/mod/golang.org/x/lint@v0.0.0-20210508222113-6edffad5e616/golint
```

# fix

> `go fix` 命令将指定代码包里面所有的 Go 标准库函数从低版本升级到高版本，自动实现兼容性。

笔者没有使用过这个命令，感兴趣的读者可以找一个第三方包的低级版本，然后尝试一下使用 `go fix` 升级。

# 格式化

## goimports

> `goimports` 命令可以对代码格式化，同时，它可以自动更新导入包的 `路径和顺序`，并且自动删除未使用的包。

```shell
# 安装 goimports
$ go install golang.org/x/tools/cmd/goimports@latest

# 格式化单个文件
$ goimports -w main.go

# 格式化整个目录
$ goimports -w foo
```

# 文档

## go doc

> `go doc` 命令可以查看包的文档。

```shell
# 查看 time.Now() 方法文档
$ go doc time.Now
# 输出如下 
package time // import "time"

func Now() Time
    Now returns the current local time.
    
    
# 查看 fmt.Printf() 方法文档
$ go doc fmt.Printf
# 输出如下
package fmt // import "fmt"

func Printf(format string, a ...any) (n int, err error)
    Printf formats according to a format specifier and writes to standard
    output. It returns the number of bytes written and any write error
    encountered.
```

除此之外，也可以直接查看某个包的文档

```shell
$ go doc strconv
# 输出省略
package strconv // import "strconv"
...
...
func Unquote(s string) (string, error)
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
type NumError struct{ ... }
```

## godoc

> godoc 用来 `提取和生成` 文档，而且还有一个很强大的功能: 将文档在浏览器内打开，实现离线文档的效果。

```shell
# 浏览器访问 Go 文档
$ godoc -http=localhost:6060
# 打开浏览器，访问 localhost:6060，可以看到 Go 所有包的文档 
```

![godoc 浏览器文档](./images/godoc_http.png)

关于 `提取文档` 的功能，限于篇幅，笔者这里不再做介绍，感兴趣的读者可以自行查阅 `godoc` 文档。

## swaggo

> Swag 通过 Go 代码注释生成标准 Swagger 文档，通过其内置的插件，可以快速与现有项目集成。

```shell
# 安装
$ go install github.com/swaggo/swag/cmd/swag@latest
# 查看版本
$ swag version 
# swag version v1.8.7
```

限于篇幅，具体的使用方法笔者这里不再做介绍，读者可以参考 [附录 2](#reference) 的 Github 主页文档，下面是一个生成的文档页面:

![swag](./images/swagger-image.png)

# lint

## golangci-lint

> golangci-lint 是一个快速的 Go linters 运行器，可以并行运行 linters 并使用缓存，支持 yaml 配置与主流 IDE 集成，包含了数十个
> linters 。

这是官网的介绍, 简单来说，`golangci-lint` 是 `静态分析工具集大成者`，包含了几乎所有的主流分析工具，
如 `errcheck`, `govet`, `statischeck`, `unused` 等，也就是说: 静态分析工具，有这一个就够了。

限于篇幅，具体的使用方法笔者这里不再做介绍，读者可以参考 [附录 4](#reference) 的官方文档。

# 小结

本小节主要介绍了常用的 Go 内置命令行工具，此外，顺带着介绍了文档生成工具 `swaggo`, 静态分析工具 `golangci-lint`,
这些都是日常开发必备工具，建议读者参照文档多加练习，做到熟练掌握。

# reference

1. [hey](https://github.com/rakyll/hey)
2. [swag](https://github.com/swaggo/swag/)
3. [golangci-lint](https://github.com/golangci/golangci-lint)
4. [golangci-lint 官网](https://golangci-lint.run/)