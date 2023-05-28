---
title: 编译文件体积优化
date: 2023-01-01
---


# 概述

通常情况下，项目编译时会通过优化来减小编译后的文件体积，这样能够加快线上服务的测试和部署流程。
接下来分别从编译选项和第三方压缩工具两方面来介绍如何有效地减小编译后的文件体积。

# 实验过程

我们以一个 `文件基础操作` 代码进行演示。

## 代码

```go
package main

import (
	"log"
	"os"
)

func fileBaseOperate(name string) (err error) {
	// 创建文件
	file, err := os.Create(name)    
	if err != nil {
		return
	}

	defer func() {
		// 关闭文件
		err = file.Close()
		if err != nil {
			return
		}
		// 删除文件
		err = os.Remove(name)
	}()

	// 向文件写入一些字符
	_, err = file.WriteString("hello world")   
	if err != nil {
		return
	}

	str := make([]byte, 1024)
	
	// 从文件读取一些字符
	_, err = file.Read(str)

	return
}

func main() {
	err := fileBaseOperate("/tmp/error_handle.log")
	if err != nil {
		log.Fatal(err)
	}
}
```

## 默认编译

```shell
$ go build main.go
$ ls -sh main
  1.9M main
```

默认编译完成的可执行文件大小是 1.9M。

## 消除符号表

默认编译完成的可执行文件会带有符号表和调试信息，发布生产时可以删除调试信息，减小可执行文件体积。

- -s：忽略符号表和调试信息。
- -w：忽略DWARFv3调试信息，使用该选项后将无法使用gdb进行调试。

```shell
$ go build -ldflags="-s -w" main.go
$ ls -sh main
  1.3M main
```

可以看到，经过 `消除符号表` 优化，编译后的文件体积已经降到了 `1.3M`, 优化了 `31%`。接下来，我们继续探索其他优化方案。

## upx

`upx` 是一个常用的压缩动态库和可执行文件的工具，通常可减少 50-70% 的文件体积。

### 安装

这里以 `MacOS` 为例，其他平台请参照 [upx Github](https://github.com/upx/upx/releases/)

```shell 
$ brew install upx
$ upx --version
  upx 3.94
```

### 使用

`upx` 有很多参数，最重要的是压缩率，`1 - 9`，1 代表最低压缩率，9 代表最高压缩率。

```shell
$ go build -ldflags="-s -w"  main.go && upx -9 main # 使用最高压缩率
$ ls -sh main
  552K main
```

可以看到，经过 `upx` 优化，编译后的文件体积已经降到了 `552KB`, 比最初的文件体积优化超过 `70%`。

### 原理

`upx` 压缩后的程序和压缩前的程序一样，无需解压仍然能够正常运行，这种压缩方法称之为带壳压缩，压缩包含两个部分：

- 在程序开头或其他合适的地方插入解压代码
- 将程序的其他部分压缩

程序执行时，也包含两个部分：

- 首先执行的是程序开头的插入的解压代码，将原来的程序在内存中解压出来
- 再执行解压后的程序，也就是说，`upx` 在程序执行时，会有额外的解压动作，不过这个耗时几乎可以忽略

# 小结

通过对示例代码的编译过程不断优化，生成的可执行文件从最开始的 `1.9M` 一直压缩到 `552K`, 压缩率超过了 `70%`，
主要是通过 **两个方法** 来实现的:

1. 编译参数 `-ldflags="-s -w"`
2. upx

# Reference

1. [upx Github](https://github.com/upx/upx/releases/)
2. [极客兔兔](https://geektutu.com/post/hpg-reduce-size.html)