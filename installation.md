# 下载
1. 打开 [Go 官网下载地址](https://go.dev/dl/)，选择对应的平台，这里以 Windows 为例
2. 点击对应的版本开始下载，比如 go1.19.1.windows-amd64.msi

# 安装
1. 双击下载好的 **.msi** 文件，然后下一步 ... 下一步 ... 最终完成

# 测试
1. 打开命令行，输入<code>go version</code>，回车，正常情况下，会输出类似下面的内容
```shell
go version go1.19.1 windows/amd64
```
2. 输入<code>go</code>，回车，正常情况下，会输出类似下面的内容
```shell
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

...
...
...

Use "go help <topic>" for more information about that topic.
```

# Hello World
和学习其他编程语言一样，写一个经典例子。

1. 打开一个目录，比如 <code>D:\Code\Go-Examples</code>
2. 新建一个文件 <code>main.go</code>，输入如下代码
```go
package main

func main() {
	println("hello world")
}
```
3. 保存文件
4. 在命令行输入 <code>go run D:\Code\Go-Examples\main.go</code>, 回车，
(当然，也可以切换到<code>D:\Code\Go-Examples</code>, 然后输入 <code>go run main.go</code>)
5. 正常情况下，会输出如下内容
```shell
hello world
```

**恭喜你，完成了 Go 的第一个程序。**

# 备注
**在后面的例子中，为了简化代码，统一默认代码路径为 <code>D:\Code\Go-Examples</code>，并且命令行已经切换完成。**