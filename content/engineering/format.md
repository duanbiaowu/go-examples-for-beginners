---
title: 格式化方法
date: 2023-01-01
---

## 格式化显示空间使用

```go
package main

import (
	"fmt"
)

func ByteCountToReadable(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func main() {
	fmt.Println(ByteCountToReadable(1024 * 1024 * 1024))
	fmt.Println(ByteCountToReadable(256 * 1024))
}
```

运行代码输出如下

```shell
$  go run main.go

1.0 GB
256.0 KB
```