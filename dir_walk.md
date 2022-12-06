# 概述

调用 `path/filepath` 包即可。`filepath.Walk()` 方法非常强大，无需递归，以非常简单的方式实现了整个目录遍历。

建议先阅读 [创建, 删除目录](dir_create_delete.md)。

# 例子

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	dir := "/tmp/test_main_go_dir"
	
	// 创建一些测试目录和文件
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		subDir := filepath.Join(dir, "1", "2", "3")
		err = os.MkdirAll(subDir, 0755) // 创建多级目录
		if err != nil {
			panic(err)
		}

		// 退出时删除目录
		defer func() {
			err = os.RemoveAll(dir)
			if err != nil {
				panic(err)
			}
		}()

		// 在目录下面新建 2 个文件
		_, err := os.Create(filepath.Join(subDir, "4.log"))
		if err != nil {
			panic(err)
		}
		_, err = os.Create(filepath.Join(subDir, "5.log"))
		if err != nil {
			panic(err)
		}
	}

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		var typ string
		if info.IsDir() {
			typ = "Dir"
		} else {
			typ = "File"
		}

		fmt.Printf("[%s] %s\n", typ, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// $ go run main.go
// 输出如下
/**
  [Dir] /tmp/test_main_go_dir
  [Dir] /tmp/test_main_go_dir/1
  [Dir] /tmp/test_main_go_dir/1/2
  [Dir] /tmp/test_main_go_dir/1/2/3
  [File] /tmp/test_main_go_dir/1/2/3/4.log
  [File] /tmp/test_main_go_dir/1/2/3/5.log
*/
```