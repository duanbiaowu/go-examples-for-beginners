## 📖 简介

帮助你快速入门 & 进阶、熟练掌握 Go 语言编程。

- [在线阅读](https://golang.dbwu.tech/)

## 🚀 快速入门

- [目录](introduction/README.md)

## 🛠️ 工程化

### 构建

- [基础开发配置](engineering/base_config.md)
- [命令工具必知必会](engineering/command.md)
- [交叉编译](engineering/compiling_cross_platform.md)
- [条件编译](engineering/conditional_compilation.md)
- [upx 优化编译文件体积](engineering/upx.md)

### 测试

- [单元测试必知必会](engineering/test.md)
- [单元测试覆盖率](engineering/test_cover.md)
- [单元测试基境](engineering/test_fixture.md)
- [基准测试数据分析](engineering/benchstat.md)
- [模糊测试-理论](engineering/test_fuzzing_theory.md)
- [模糊测试-实践](engineering/test_fuzzing_practice.md)
- [压力测试](engineering/test_performance.md)

### 实践

- [channel 操作规则](engineering/channel.md)
- [结构体使用技巧](engineering/struct.md)
- [切片使用技巧](engineering/slice.md)
- [JSON 使用技巧](engineering/json.md)
- [embed 嵌入文件](engineering/embed.md)
- [expvar 监控接口状态](engineering/expvar.md)
- [Go 的面向对象编程](engineering/oop_in_go.md)
- [如何实现 implements](engineering/implements.md)
- [数据竞态](engineering/data_race.md)
- [错误处理最佳实践](engineering/error_handle_gracefully.md)
- [Gin 快速入门](engineering/gin)
- [zap 快速入门](engineering/zap)
- [wire 快速入门](engineering/wire)
- [保留小数位数](engineering/util/math.go)
- [格式化显示占用空间](engineering/util/file.go)

## ☹️ 陷阱

- [数组和切片参数传递差异](traps/array_with_map_in_params.md)
- [byte 加减](traps/byte_operation.md)
- [map](traps/map_struct_assign.md)
- [copy 复制失败](traps/copy.md)
- [缓冲区内容不输出](traps/buffer_flush.md)
- [切片占用过多内存](traps/slice_occupy_memory.md)
- [实现 String 方法陷入无限递归](traps/string_method.md)
- [错误处理三剑客](traps/defer_with_recover.md)
- [几个有趣的 defer 笔试题](traps/defer_exam.md)
- [nil != nil ?](traps/nil_with_nil.md)
- [nil 作为参数引发的问题](traps/nil_argument.md)
- [for 循环赋值错误](traps/for_assign.md)
- [for 循环调用函数](traps/for_func.md)
- [for 循环 goroutine 执行顺序不一致](traps/for_goroutine.md)
- [interface 方法调用规则](traps/interface_method.md)
- [interface{} != *interface{} ?](traps/interface_error.md)
- [goroutine 竞态](traps/goroutine_race.md)
- [goroutine 泄漏](traps/channel_not_closed.md)

## ⚡ 高性能

- [for](performance/for.md)
- [切片预分配](performance/slice_pre_alloc.md)
- [切片过滤器](performance/slice_filter.md)
- [切片和数组](performance/slice_with_array.md)
- [string 与 []byte 转换](performance/string_with_bytes.md)
- [map 预分配](performance/map_pre_alloc.md)
- [map key 类型](performance/map_key_type.md)
- [map 重置和删除](performance/map_free.md)
- [整数转字符串](performance/int_to_string.md)
- [字符串拼接](performance/string_concat.md)
- [截取中文字符串](performance/sub_cn_string.md)
- [空结构体](performance/empty_struct.md)
- [结构体切片](performance/struct_slice.md)
- [对象复用](performance/sync_pool.md)
- [获取调用堆栈优化](performance/stack_dump.md)
- [字节序优化](performance/binary_read_write.md)
- [goroutine ID](performance/goroutineid.md)
- [defer 优化](performance/defer.md)
- [timer 优化](performance/timer.md)
- [channel 缓冲和非缓冲](performance/channel.md)
- [互斥锁和读写锁](performance/mutex.md)
- [内联优化](performance/inline.md)
- [内存对齐](performance/memory_alignment.md)
- [逃逸分析](performance/escape.md)
- [singleflight](performance/singleflight.md)

## 📚 设计模式

### 创建型模式

- [构建](patterns/creational/builder.go)
- [工厂](patterns/creational/factory.go)
- [对象池](patterns/creational/object_pool.go)
- [单例](patterns/creational/singleton.go)

### 结构性模式

- [适配器](patterns/structural/adapter.go)
- [装饰者](patterns/structural/decorator.go)
- [代理](patterns/structural/proxy.go)

### 行为型模式

- [责任链](patterns/behavioral/chain_of_responsibility.go)
- [观察者](patterns/behavioral/observer.go)
- [状态](patterns/behavioral/state.go)
- [策略](patterns/behavioral/strategy.go)

### 其他模式

- [并发模式](patterns/concurrency/)
- [常用模式](patterns/idiom/)
- [K8S](patterns/k8s/visitor.go)
- [MapReduce](patterns/mapreduce/real_world.go)

## 微信

![微信公众号](introduction/images/wechat_accounts.png)

## JetBrains open source certificate support

This project has always been developed in the GoLand integrated development environment under JetBrains, based on the **free JetBrains Open Source license(s)** genuine free license. I would like to express my gratitude.

<a href="https://www.jetbrains.com/?from=Duan Biaowu" target="_blank"><img src="introduction/images/jetbrain.png" style="width: 36%;"/></a>