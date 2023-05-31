---
headless: true
---


### [Blog](https://dbwu.tech/)
### [Github](https://github.com/duanbiaowu/go-examples-for-beginners)
### [微信公众号](https://dbwu.tech/images/wechat.png)

---

### 🛠️ 工程化

- **构建**
  - [开发环境配置](../engineering/base_config.md)
  - [命令行工具](../engineering/command.md)
  - [交叉编译](../engineering/compiling_cross_platform.md)
  - [条件编译](../engineering/conditional_compilation.md)
  - [编译文件体积优化](../engineering/upx.md)

- **测试**
  - [单元测试基础必备](../engineering/test.md)
  - [单元测试覆盖率](../engineering/test_cover.md)
  - [单元测试基境](../engineering/test_fixture.md)
  - [基准测试数据分析](../engineering/benchstat.md)
  - [模糊测试-理论](../engineering/test_fuzzing_theory.md)
  - [模糊测试-实践](../engineering/test_fuzzing_practice.md)
  - [压力测试](../engineering/test_performance.md)

- **实践**
  - [Go 的面向对象编程](https://dbwu.tech/posts/golang_oop/)
  - [如何实现 implements](https://dbwu.tech/posts/golang_implements/)
  - [channel 操作规则](../engineering/channel.md)
  - [结构体使用技巧](../engineering/struct.md)
  - [切片使用技巧](../engineering/slice.md)
  - [JSON 使用技巧](../engineering/json.md)
  - [embed 嵌入文件](../engineering/embed.md)
  - [expvar 监控接口状态](../engineering/expvar.md)
  - [数据竞态](https://dbwu.tech/posts/golang_data_race/)
  - [错误处理最佳实践](../engineering/error_handle_gracefully.md)
  - [Gin 快速入门](https://github.com/duanbiaowu/go-examples-for-beginners/blob/master/engineering/gin/quick_start.go)
  - [zap 快速入门](https://github.com/duanbiaowu/go-examples-for-beginners/blob/master/engineering/zap/quickstart_test.go)
  - [wire 快速入门](https://github.com/duanbiaowu/go-examples-for-beginners/blob/master/engineering/wire/README.md)
  - [常用数学方法](../engineering/math.md)
  - [格式化方法](../engineering/format.md)

---

### ☹️ 陷阱

- **仔细检查你的代码**
  - [数组和切片参数传递差异](../traps/array_with_map_in_params.md)
  - [byte 加减](../traps/byte_operation.md)
  - [map 常见问题](../traps/map_struct_assign.md)
  - [copy 函数复制失败](../traps/copy.md)
  - [缓冲区内容不输出](../traps/buffer_flush.md)
  - [切片占用过多内存](../traps/slice_occupy_memory.md)
  - [String 方法陷入无限递归](../traps/string_method.md)
  - [错误处理三剑客](../traps/defer_with_recover.md)
  - [几个有趣的 defer 笔试题](../traps/defer_exam.md)
  - [nil != nil ?](../traps/nil_with_nil.md)
  - [nil 作为参数引发的问题](../traps/nil_argument.md)
  - [for 循环赋值错误](../traps/for_assign.md)
  - [for 循环调用函数](../traps/for_func.md)
  - [for 循环 goroutine 执行顺序不一致](../traps/for_goroutine.md)
  - [interface 方法调用规则](../traps/interface_method.md)
  - [interface{} != *interface{} ?](../traps/interface_error.md)
  - [goroutine 竞态](../traps/goroutine_race.md)
  - [goroutine 泄漏](../traps/channel_not_closed.md)

---

### ⚡ 高性能

- **让你的代码速度飞起**
  - [for](../performance/for.md)
  - [切片预分配](../performance/slice_pre_alloc.md)
  - [切片过滤器](../performance/slice_filter.md)
  - [切片和数组](../performance/slice_with_array.md)
  - [string 与 []byte 转换](../performance/string_with_bytes.md)
  - [map 预分配](../performance/map_pre_alloc.md)
  - [map key 类型](../performance/map_key_type.md)
  - [map 重置和删除](../performance/map_free.md)
  - [整数转字符串](../performance/int_to_string.md)
  - [字符串拼接](../performance/string_concat.md)
  - [截取中文字符串](../performance/sub_cn_string.md)
  - [空结构体](../performance/empty_struct.md)
  - [结构体切片](../performance/struct_slice.md)
  - [对象复用](../performance/sync_pool.md)
  - [获取调用堆栈优化](../performance/stack_dump.md)
  - [字节序优化](../performance/binary_read_write.md)
  - [获取 goroutine ID](../performance/goroutineid.md)
  - [defer 优化](../performance/defer.md)
  - [timer 优化](../performance/timer.md)
  - [channel 缓冲和非缓冲](../performance/channel.md)
  - [互斥锁和读写锁](../performance/mutex.md)
  - [内联优化](../performance/inline.md)
  - [内存对齐](../performance/memory_alignment.md)
  - [逃逸分析](https://dbwu.tech/posts/goland_escape/)
  - [singleflight](../performance/singleflight.md)