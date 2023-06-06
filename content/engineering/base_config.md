---
title: 开发环境配置
date: 2023-01-01
---

# 概述

本小节主要讲述如何快速搭建一个现代化的 Go 开发环境。

# 基础环境变量

## GOROOT

Go 源代码的安装目录，`Mac` 和 `Windows`  安装时会自动配置好，`Linux` 一般在 `/usr/local/go` 目录。

```shell
# 查看 $GOROOT 目录
$ echo $GOROOT
/usr/local/go

# 设置 $GOROOT 目录
$ export GOROOT=/usr/local/go

# 增加 go 相关命令到 PATH
$ export PATH=$PATH:/usr/local/go/bin
```

## GOPATH

简单来说，就是存放 Go 第三方库的源代码以及构建后可执行程序的目录，建议设置为独立的目录并且不要存放其他文件。

```shell
# 查看 $GOPATH 目录
$ echo $GOPATH
/home/codes/go

# 设置 $GOPATH 目录
$ export GOPATH=/home/codes/go
```

## GOPROXY

安装包的下载代理地址，直接使用 `七牛云` 提供的代理地址 (https://goproxy.cn)，速度非常快！

```shell
# 查看 $GOPROXY 代理地址 
$ echo $GOPROXY
https://goproxy.cn

# 设置 $GOPROXY 目录
$ export GOPROXY=https://goproxy.cn
```

## GO111MODULE

是否开启了 `gomod`, 必须开启才可以使用 `Go Module` 。

```shell
# 查看 $GO111MODULE 模块开启情况 
$ echo $GO111MODULE
on

# 设置 $GOPROXY 目录
$ export GO111MODULE="on"
```

**建议将刚才的配置写入 `~/.bashrc` 或者 `~/.zshrc` 文件，永久有效。**
到这里，我们第一步设置基础环境变量的工作就完成了，可以验证一下设置是否成功:

```shell
$ go env | grep -i -E "root|path|goproxy|module"
GO111MODULE="on"
GOPATH="/home/codes/go"
GOPROXY="https://goproxy.cn"
GOROOT="/usr/local/go"
```

通过输出结果可以看到，刚才的配置已经全部完成。

## Go 命令方法

`env` 命令除了查看环境变量外，同样可以设置环境变量。

```shell
# 通过 go env 命令设置代理环境变量
$ go env -w GOPROXY=https://goproxy.cn,direct
```

# Modules

Go 从 1.12 版本开始，默认支持 `Go Modules`, 从此彻底告别配置 `GOPATH` 以及包下载及依赖导致的各种奇葩问题。

## 初始化一个包

这里假设项目名称为 `HelloWorld` 。

- 新建项目的目录，比如 `/home/codes/projects/HelloWorld`
- 切换到 `/home/codes/projects/HelloWorld` 目录
- 执行命令 `go mod init helloworld`

```shell
# 输出如下
go: creating new go.mod: module helloworld
go: to add module requirements and sums:
go mod tidy
```

- 这时可以看到目录下多了一个 `go.mod` 文件，其中内容如下

```shell
$ cat go.mod
module helloworld

go 1.19  # 版本号可能和你的不一样
```

## 安装依赖

Go 的包名定义非常简单，就是一个普通的 URL (以域名打头)，可以是主流的代码仓库地址，也可以是自己搭建的代码仓库。
下面的例子统一以 Github 演示。

```shell
# 语法规则: 其中版本号可以是 git 分支或 tag
go get 包名@版本号
# 例: go get github.com/spf13/cast@v1.4.1
```

### 安装 spf13/cast 包

`spf13/cast` 是一个数据类型转换包，可以非常简单地对常见数据类型互相转换，并且不会引发 `panic` 。

- 执行命令:

```shell
$ go get github.com/spf13/cast@v1.4.1
# 输出如下
go: added github.com/spf13/cast v1.4.1
```

- 打开 `go.mod` 文件，内容如下:

```shell
module helloworld

go 1.19

require github.com/spf13/cast v1.4.1 // indirect
```

- 在 `go.mod` 旁边多了一个 `go.sum` 文件，内容如下:

```shell
github.com/davecgh/go-spew 
...
... 
github.com/stretchr/testify 
```

- 使用安装好的包

将如下代码写入文件 `main.go`

```go
package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	s := cast.ToString(1024)
	fmt.Printf("s is a %T, val = %s\n", s, s)
}
```

```shell
$ go run main.go
// 输出如下
/**
  s is a string, val = 1024
*/
```

## 查看依赖

```shell
# 列表输出
$ go list -m -m all
# # json 输出
$ go list -m -json all 

# 输出当前项目的 Module 名称以及依赖报名
helloworld
...
...
github.com/spf13/cast v1.4.1
... 
... 
```

## 升级依赖

```shell
# 语法规则: 
# 升级次级或补丁版本号
go get -u 包名@版本号
# 仅升级补丁版本号
go get -u=patch 包名@版本号
```

这里，我们将 `spf13/cast` 包从 `v1.4.1` 升级到 `1.5.0`，执行如下命令:

```shell
go get -u github.com/spf13/cast@v1.5.0
# 输出如下
go: upgraded github.com/spf13/cast v1.4.1 => v1.5.0
```

查看 `go.mod` 文件，内容已经更新为:

```shell
module helloworld

go 1.19

require github.com/spf13/cast v1.5.0 // indirect
```

## 删除依赖

当前项目中有些包已经不再使用了，但是 `go.mod` 文件中依然定义了依赖关系，可以使用下面的命令**自动整理优化** `go.mod` 文件。

```shell
$ go mod tidy
```

## 常用命令

```shell
go mod init  # 初始化
go mod tidy  # 更新 (移除) 依赖文件
go mod download  # 下载依赖文件

go mod vendor  # 将依赖全部归档到 vendor 目录
go mod edit    # 修改依赖文件
go mod graph   # 打印依赖关系 图
go mod verify  # 校验依赖
```

到这里，`Go Modules` 的基础配置及使用已经完成，我们可以快速导入成熟的第三方库来加速开发。

# 编辑器

**工欲善其事，必先利其器。** 现代化项目开发，一个高效的 IDE 必不可少，下面引入几篇编辑器相关的文档 (实在写不动了 ^_^)。

## Goland

- [安装及配置](https://polarisxu.studygolang.com/posts/go/2022-dev-env/)
- [如何获取免费的授权](https://strikefreedom.top/archives/jetbrains-open-source-license-for-free)

## VSCode

- [安装及配置](https://mp.weixin.qq.com/s/J01LY7s6xMB8Lk10sxTFhg)
- [安装及配置 - 2](https://colobu.com/2016/04/21/use-vscode-to-develop-go-programs)

# 小结

这篇文章主要围绕 `Go 环境变量`, `Go Modules`, `编辑器` 三个方面展开讲述，环境变量是基础配置以及 Modules 的基础，
Module 是对包进行下载和更新等操作的管理工具，最后引入了几篇不错的文档讲解如何安装和配置主流的编辑器。
到这里，读者应该搭建好自己的开发环境了，还在等什么呢？赶紧去开发一个自己的项目吧 :)
