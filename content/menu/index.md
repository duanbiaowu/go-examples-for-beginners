---
headless: true
---


- [Blog](https://dbwu.tech/)
- [Github](https://github.com/duanbiaowu/go-examples-for-beginners)
- [微信公众号](https://dbwu.tech/images/wechat.png)

---

## 快速入门系列

- [**前言**](../introduction/preface.md)


- **1. 环境安装**
  - [Windows](../introduction/windows.md)
  - [Mac](../introduction/mac.md)
  - [Linux](../introduction/linux.md)
  - [合并版](../introduction/install.md)

---

- **2. 语法基础**
  - [Go 程序的运行方式及要求](../introduction/buildandrun.md)
  - [基本代码格式，关键字, 预定义标识符](../introduction/basesyntaxandkeyword.md)
  - [变量](../introduction/variables.md)
  - [空白标识符](../introduction/blank_operator.md)
  - [常量](../introduction/const.md)
  - [自定义类型](../introduction/typedef.md)
  - [常量生成器](../introduction/iota.md)
  - [运算优先级](../introduction/operator_priority.md)
  - [包的导入](../introduction/import.md)
  - [调试打印](../introduction/print.md)
  - [字符串](../introduction/string.md)
  - [字符](../introduction/rune.md)
  - [类型转换](../introduction/data_convert.md)
  - [保留小数位](../introduction/decimal.md)
  - [指针](../introduction/pointer.md)
  - [if/else](../introduction/if_else.md)
  - [自增自减](../introduction/inc_and_dec.md)
  - [for](../introduction/for.md)
  - [range](../introduction/range.md)
  - [switch](../introduction/switch.md)
  - [goto / 标签](../introduction/goto.md)
  - [可见性](../introduction/visable.md)
  - [作用域](../introduction/scope.md)

---

- **3. 数据类型**
  - [数组](../introduction/array.md)
  - [切片](../introduction/slice.md)
  - [字符切片](../introduction/bytes.md)
  - [Map](../introduction/map.md)
  - [有序 Map](../introduction/sorted_map.md)
  - [函数](../introduction/func.md)
  - [init](../introduction/init.md)
  - [make, new](../introduction/make_with_new.md)
  - [变长参数](../introduction/func_variadic_params.md)
  - [指针参数](../introduction/func_pointer_params.md)
  - [闭包](../introduction/func_closures.md)
  - [递归](../introduction/func_recursion.md)
  - [内部函数](../introduction/func_inner.md)
  - [panic](../introduction/panic.md)
  - [defer](../introduction/defer.md)
  - [recover](../introduction/recover.md)
  - [结构体](../introduction/struct.md)
  - [嵌套结构体](../introduction/struct_embedding.md)
  - [方法](../introduction/methods.md)
  - [接口](../introduction/interface.md)
  - [实现系统错误接口](../introduction/implement_error.md)
  - [判断是否实现接口](../introduction/implement.md)
  - [错误](../introduction/error.md)
  - [零值](../introduction/zero_value.md)
  - [类型比较](../introduction/type_comparison.md)

---

- **4. 协程与通道**
  - [goroutine](../introduction/goroutine.md)
  - [非缓冲通道](../introduction/channel.md)
  - [缓冲通道](../introduction/channel_buffer.md)
  - [关闭通道](../introduction/channel_close.md)
  - [通道方向](../introduction/channel_direction.md)
  - [检测通道是否关闭](../introduction/channel_close_check.md)
  - [遍历通道](../introduction/channel_range.md)
  - [waitgroup](../introduction/waitgroup.md)
  - [select](../introduction/select.md)
  - [互斥锁](../introduction/mutex.md)
  - [超时控制](../introduction/timeout.md)
  - [定时器](../introduction/ticker.md)

---

- **5. 常见操作**
  - [原子操作](../introduction/atomic.md)
  - [创建, 删除文件](../introduction/file_create_delete.md)
  - [写文件](../introduction/file_write.md)
  - [读文件](../introduction/file_read.md)
  - [文件路径, 扩展名](../introduction/file_path.md)
  - [文件判断](../introduction/file_check.md)
  - [创建, 删除目录](../introduction/dir_create_delete.md)
  - [遍历目录](../introduction/dir_walk.md)
  - [日志](../introduction/log.md)
  - [HTTP](../introduction/http.md)
  - [URL](../introduction/url.md)
  - [base64](../introduction/base64.md)
  - [sha256](../introduction/sha256.md)
  - [md5](../introduction/md5.md)
  - [exit](../introduction/exit.md)
  - [获取进程ID](../introduction/process_id.md)
  - [命令行](../introduction/command.md)
  - [命令行参数](../introduction/command_args.md)
  - [命令行参数解析与设置](../introduction/command_flag.md)
  - [信号](../introduction/signal.md)
  - [json](../introduction/json.md)
  - [xml](../introduction/xml.md)
  - [日期, 时间](../introduction/time.md)
  - [时间戳](../introduction/timestamp.md)
  - [随机数](../introduction/random.md)
  - [正则表达式](../introduction/regexp.md)

---

## 进阶提升系列

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

---

### <a href="https://dbwu.tech/posts/golang_oop/" target="_blank">Go 面向对象编程</a>

---

### 📚 设计模式

- **创建型模式**
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/creational/builder.go" target="_blank">构建</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/creational/factory.go" target="_blank">工厂</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/creational/object_pool.go" target="_blank">对象池</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/creational/singleton.go" target="_blank">单例</a>

- **结构性模式**
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/structural/adapter.go" target="_blank">适配器</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/structural/decorator.go" target="_blank">装饰器</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/structural/proxy.go" target="_blank">代理</a>

- **行为型模式**
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/behavioral/chain_of_responsibility.go" target="_blank">责任链</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/behavioral/observer.go" target="_blank">观察者</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/behavioral/state.go" target="_blank">状态</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/behavioral/strategy.go" target="_blank">策略</a>

- **其他模式**
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/concurrency/" target="_blank">并发模式</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/idiom/" target="_blank">常用模式</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/k8s/visitor.go" target="_blank">K8S</a>
  - <a href="https://github.com/duanbiaowu/go-examples-for-beginners/tree/master/patterns/mapreduce/real_world.go" target="_blank">MapReduce</a>