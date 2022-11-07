# 概述

## 两者的区别
- new(T) 为数据类型 T 分配一块内存，初始化为类型 T 的零值，返回类型为指向数据内存的指针，可以用于所有数据类型。
- make(T) 除了为数据类型 T 分配内存外，还可以指定长度和容量，返回类型为数据的初始化结构，只限于 `切片`, `Map`, `通道`。

<p align="center">
<img width="600" src="./images/make_with_new.jpg">
</p>

## 什么时候用 make()?
[切片](slice.md), [Map](map.md), [通道](channel.md)

## 什么时候用 new()?
[数组](array.md), [结构体](struct.md) 和所有其他类型。
