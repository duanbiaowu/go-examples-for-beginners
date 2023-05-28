---
title: 交叉编译
date: 2023-01-01
---


# 概述

交叉编译，也称跨平台编译，就是在一个平台上编译源代码，生成结果为另一个平台上的可执行代码。
这里的平台包含两个概念：体系架构 (如 AMD, ARM) 和 操作系统 (如 Linux, Windows）。
同一个体系架构可以运行不同的操作系统，反过来，同一个操作系统也可以运行在不同的体系架构上。

> Go 实现跨平台编译的思想其实很简单：通过保存可以生成最终机器码的多份翻译代码，
> 在编译时根据 GOARCH=体系架构 和GOOS=操作系统参数进行初始化设置，
> 最终调用对应平台编写的特定方法来生成机器码，从而实现跨平台编译。

# 例子

下面的例子统一以 `amd64` 作为体系架构参数，读者请根据自己的环境更换对应参数，比如 `386`。

## Mac

### 编译为 Linux 代码

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# 也可以是 386 平台
# CGO_ENABLED=0 GOOS=linux GOARCH=386 go build main.go
```

### 编译为 Windows 代码

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

```

## Linux

### 编译为 Mac 代码

```shell
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

### 编译为 Windows 代码

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## Windows

### 编译为 Mac 代码

```shell
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go
```

### 编译为 Linux 代码

```shell
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

# 扩展阅读

1. [交叉编译 - 维基百科](https://zh.wikipedia.org/wiki/%E4%BA%A4%E5%8F%89%E7%B7%A8%E8%AD%AF%E5%99%A8)
2. [含有CGO代码的项目如何实现跨平台编译](https://segmentfault.com/a/1190000038938300)