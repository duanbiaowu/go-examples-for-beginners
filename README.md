# 目录

前言
  为什么要写这本书
    总结输出
    提炼笔记
    锻炼写作水平
    体会自己了解的东西，教会别人的过程。费曼学习法的践行。
    体验分享的感觉
    认识新朋友
  建议顺序阅读

- [环境安装](installation.md)
- Windows
- Mac
- Linux
- [Go 程序的运行方式及要求](buildandrun.md)
- [基本代码格式，关键字, 预定义标识符](basesyntaxandkeyword.md)
- [变量](variables.md)
- [空白标识符](blank_operator.md)
- [常量](const.md)
- [自定义类型](typedef.md)
- [常量生成器](iota.md)
- [运算优先级](operator_priority.md)
- 字符
- 字符串 中文，编码

- [包的导入](import.md)
- [调试打印](print.md)
- [自增自减](inc_and_dec.md)
- [类型转换](data_convert.md)
- [保留小数位](decimal.md)
- [指针](pointer.md)
- uintptr 
- nil
[rune](rune.md)
[错误](error.md)
- 零值

- [if/else](if_else.md)
- [for](for.md)
- [range](range.md)
- [switch](switch.md)
- [goto / 标签](goto.md)
可见性 (大写，小写)
作用域

- [数组](array.md)
- [切片](slice.md)
- [Map](map.md)
- [有序 Map](sorted_map.md)

- [函数](func.md)
- [init](init.md)
[make, new](make_with_new.md)
[可变参数](func_variadic_params.md)
[指针参数](func_pointer_params.md)
[闭包](func_closures.md)
[递归](func_recursion.md)
[内部函数](func_inner.md)

[panic](panic.md)
[defer](defer.md) 
[recover](recover.md)

[结构体](struct.md)
[嵌套结构体](struct_embedding.md)
[方法](methods.md)
[接口](interface.md)
[实现系统错误接口](implement_error.md)
[判断是否实现接口](implement.md)

[goroutine](goroutine.md)
[waitgroup](waitgroup.md)
[阻塞通道](channel.md)
[非阻塞通道](channel_buffer.md)
[关闭通道](channel_close.md)
[通道方向](channel_direction.md)
[检测通道是否关闭](channel_close_check.md)
[遍历通道](channel_range.md)
[select](select.md)

- [进程ID](process_id.md)

[json](json.md)
[xml](xml.md)
[日期, 时间](time.md)
[时间戳](timestamp.md)
[random](random.md)
[正则表达式](regexp.md)

- [创建, 删除文件](file_create_delete.md)
- [写文件](file_write.md)
- [读文件](file_read.md)
- [文件路径, 扩展名](file_path.md)
- [文件判断](file_check.md)
- [创建, 删除目录](dir_create_delete.md)
- [遍历目录](dir_walk.md) 

- [exit](exit.md)
- [命令行](command.md)
- [命令行参数](command_args.md)
- [命令行参数解析与设置](command_flag.md)
- [信号](signal.md)

- [日志](log.md)
- [HTTP](http.md) 
- [URL](url.md)
- [base64](base64.md)
- [sha256](sha256.md)
- [md5](md5.md)

- [超时控制](timeout.md)
- [定时器](ticker.md)

- [原子操作](atomic.md)
[互斥锁](mutex.md)

泛型