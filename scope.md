# 概述
`词法块` 是指由大括号围起来的一个语句序列，比如 `for` 循环语句块，`if/else` 判断语句块。
在 `语句块` 内部声明的变量对外部不可见，块把声明围起来，决定了它的作用域。

**一个程序可以包含多个相同名称的变量，只要这些变量处在不同的 `语句块` 内。** 
虽然语法上支持，但是实践中不建议这样做。

# 全局变量
全局变量根据 [可见行](visable.md) 规则，决定在包内可用还是全局可用。

# 导入的包

# 小结
本小结所有的声明均以变量为例子，所有规则同样适用于常量以及其他声明。

# reference
1. https://book.douban.com/subject/27044219/