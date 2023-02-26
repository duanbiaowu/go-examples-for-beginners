# 概述

为了节省篇幅，笔者将常用的 3 种操作系统对应的安装教程汇总到了一起，读者可以直接选择对应内容阅读。

# Windows 环境搭建

## 下载
1. 打开 Go 官网下载地址(https://go.dev/dl/)，选择 `Microsoft Windows`
2. 点击对应的版本开始下载，比如 `go1.19.1.windows-amd64.msi`

## 安装

双击下载好的 **.msi** 文件，然后下一步 -> 下一步 -> 最终完成

## 测试

打开命令行，输入`go version`，回车，正常情况下，会输出类似下面的内容

```shell
go version go1.19.1 windows/amd64
```

输入 `go`，回车，正常情况下，会输出类似下面的内容

```shell
Go is a tool for managing Go source code.
Usage:
  go <command> [arguments]
The commands are:  
...
...
...
Use "go help <topic>" for more information about that topic.
```

## Hello World
    
和学习其他编程语言一样，写一个经典例子。

- 打开一个目录，比如 `D:\Code\Go-Examples`
- 新建一个文件 `main.go`，输入如下代码
    
```go
package main

func main() {
    println("hello world")
}
```

- 保存文件 
- 在命令行输入 `go run D:\Code\Go-Examples\main.go`, 回车，(当然，也可以切换到`D:\Code\Go-Examples`, 然后输入 `go run main.go`)
- 正常情况下，会输出如下内容

```shell
    hello world
```

**恭喜你，完成了 Go 的第一个程序。**

## 备注
**在后面的例子中，为了简化代码，统一默认代码路径为 `D:\Code\Go-Examples`，并且目录已经切换完成。**

# Mac 安装

## 下载
1. 打开 Go 官网下载地址(https://go.dev/dl/)
2. 根据硬件架构选择 `Apple macOS (ARM64)` 或 `Apple macOS (x86-64)`
3. 点击对应的版本开始下载，比如 `go1.19.1.darwin-arm64.pkg`

## 安装

双击下载好的 **.pkg** 文件，后续过程和安装其他 Mac App 一样

## 测试
 
打开命令行，输入`go version`，回车，正常情况下，会输出类似下面的内容

```shell
go version go1.19.1 darwin/arm64
```
   
- 输入`go`，回车，正常情况下，会输出类似下面的内容

```shell
Go is a tool for managing Go source code.
Usage:
  go <command> [arguments]
The commands are:
...
...
...
Use "go help <topic>" for more information about that topic.
```

## Hello World

和学习其他编程语言一样，写一个经典例子。

- 打开一个目录，比如 `/Users/codes/Go-Examples`
- 新建一个文件 `main.go`，输入如下代码

```go
package main

func main() {
    println("hello world")
}
```

- 保存文件 
- 在命令行输入 `go run /Users/codes/Go-Examples/main.go`, 回车，
   (当然，也可以切换到`/Users/codes/Go-Examples`, 然后输入 `go run main.go`)
- 正常情况下，会输出如下内容

```shell
hello world
```

**恭喜你，完成了 Go 的第一个程序。**

### 备注

**在后面的例子中，为了简化代码，统一默认代码路径为 `/Users/codes/Go-Examples`，并且目录已经切换完成。**

# Linux 安装

## 下载
1. 打开 Go 官网下载地址(https://go.dev/dl/)
2. 根据硬件架构选择 `Linux` (已编译完成) 或 `Source` (源代码)，这里以编译完的发行版为例
3. 点击对应的版本压缩包开始下载，比如 `go1.19.1.linux-amd64.tar.gz`

## 安装

直接将压缩包文件解压到 `/usr/local/` 目录

```shell
sudo tar -zxvf go1.19.1.linux-amd64.tar.gz -C /usr/local/
```

## 测试

打开命令行，输入`go version`，回车，正常情况下，会输出类似下面的内容
   
```shell
go version go1.19.1 linux/amd64
```
   
输入`go`，回车，正常情况下，会输出类似下面的内容

```shell
Go is a tool for managing Go source code.
Usage:
  go <command> [arguments]  
...
...
...
Use "go help <topic>" for more information about that topic.
```

## Hello World

和学习其他编程语言一样，写一个经典例子。

- 打开一个目录，比如 `/home/codes/Go-Examples`
- 新建一个文件 `main.go`，输入如下代码
    
```go
package main

func main() {
    println("hello world")
}
```
   
- 保存文件 
- 在命令行输入 `go run /home/codes/Go-Examples/main.go`, 回车，
(当然，也可以切换到`/home/codes/Go-Examples`, 然后输入 `go run main.go`)
- 正常情况下，会输出如下内容

```shell
hello world
```

**恭喜你，完成了 Go 的第一个程序。**

### 备注
**在后面的例子中，为了简化代码，统一默认代码路径为 `/home/codes/Go-Examples`，并且目录已经切换完成。**