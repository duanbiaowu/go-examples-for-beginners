## 箴言

- 异步调用的选择权交给调用方
    - 因为调用方可能并不知道函数内部实现使用了 goroutine
- 准备启动一个 goroutine 时
    - 永远不要启动无法控制退出的 goroutine
    - 永远不要启动无法确定何时退出的 goroutine
    - 启动 goroutine 时实现 `panic recovery` 机制，避免服务内部错误导致的不可用
    - 造成 goroutine 泄漏的主要原因是 goroutine 阻塞，并且无法控制其退出
- 避免在请求中直接启动 goroutine
    - 应该通过启动 `worker` 消费者模式来处理，可以避免由于请求量过大，创建大量 goroutine 导致的 `OOM`
- goroutine 只能自己退出，而不能被其他 goroutine 强制关闭或杀死
    - goroutine 被设计成不可以从外部无条件地退出，只能通过 `channel` 进行通信退出
    - 杀死一个 goroutine 设计上会有很多挑战，当前所拥有的资源如何处理？堆栈如何处理？`defer` 语句需要执行么？
    - 如果允许 defer 语句执行，那么 `defer` 语句可能阻塞 goroutine 退出，这种情况下怎么办呢？
- 使用建议
    - 尽量使用 `非阻塞 I/O`（非阻塞 I/O 常用来实现高性能的网络库），`阻塞 I/O` 很可能导致 goroutine 在某个调用一直等待，而无法正确结束
    - 任务分段执行，超时后立即退出，避免 goroutine 浪费资源

## reference

1. https://lailin.xyz/post/go-training-week3-goroutine.html