## 两个重要的概念

### provider

是一个普通的函数，会返回构建依赖关系所需的组件。实际项目中，往往是一些简单的工厂函数，不会太复杂。

**注意：不能存在两个 provider 返回相同的组件类型。**

```go
type Player struct {
    Name string
}

func NewPlayer(name string) Player {
    return Player{name}
}
```

### injector

是一个对象的创建和初始化函数。

```shell
# 详情见 demo
./guide/advanced/struct_providers/
```

## Reference

- [google/wire](https://github.com/google/wire)
- [Go 每日一库之 wire](https://zhuanlan.zhihu.com/p/110453784)