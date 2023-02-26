# 概述

调用 `time` 包即可。

# 例子

## 获取/解析 时间戳

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Unix())      // 单位: 秒
	fmt.Println(now.UnixMilli()) // 单位: 毫秒
	fmt.Println(now.UnixMicro()) // 单位: 微妙
	fmt.Println(now.UnixNano())  // 单位: 纳秒

	var timestamp int64 = 1557433059
	t := time.Unix(timestamp, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

// $ go run main.go
// 输出如下，你的输出可能和这里的不一样
/**
  1667433164
  1667433164630
  1667433164630949
  1667433164630949000
  2019-05-10 04:17:39
*/
```