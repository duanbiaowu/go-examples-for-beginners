# 原生 log

## 优点

简单、开箱即用，不用引用外部的三方库。

## 缺点

* 仅限基本的日志级别
    * 只有一个 Print 选项。不支持 INFO/DEBUG 等多个级别。
* 对于错误日志，它有 Fatal 和Panic
    * Fatal 日志通过调用os.Exit(1)来结束程序
    * Panic 日志在写入日志消息之后抛出一个panic
    * 缺少 ERROR 日志级别，这个级别可以在不抛出 panic 或退出程序的情况下记录错误
* 缺乏结构化日志格式的能力——只支持简单文本输出，不支持 JSON 等自定义格式。
* 不提供日志切割的能力。

# zap

## 优点

* 高性能
* 优秀的设计与工程实践

## 高性能实现

* 避免 GC: 对象复用
* 避免反射: 内建 Encoder
* 避免竞态
* 避免加锁, 写时复制
* 避免使用 interface{} 带来的开销（拆装箱、对象逃逸到堆上）
* 避免使用 fmt json/encode 使用字符编码方式对日志信息编码，适用 []byte 

# Reference

1. https://segmentfault.com/a/1190000022461706
2. https://mp.weixin.qq.com/s/Zif7HnNV1y7swunNQzZopQ
3. https://mp.weixin.qq.com/s/i0bMh_gLLrdnhAEWlF-xDw
4. https://medium.com/a-journey-with-go/go-how-zap-package-is-optimized-dbf72ef48f2d