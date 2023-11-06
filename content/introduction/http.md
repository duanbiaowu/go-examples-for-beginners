# 概述

`net/http` 包含了 HTTP 相关方法。

# 例子

```go
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Printf("resp.Body.Close() %s", err)
		}
	}()

	fmt.Printf("Response status code = %d\n", resp.StatusCode)
	fmt.Printf("Response content type = %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("Response content length = %d\n", resp.ContentLength)

	body := make([]byte, resp.ContentLength)
	n, err := resp.Body.Read(body)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	fmt.Printf("Response body read = %d\n", n)
	fmt.Printf("Response body = %s\n", body)
}

// $ go run main.go
// 输出如下
/**
  Response status code = 200
  Response content type = text/html
  Response content length = 227
  Response body read = 227
  Response body = <html>
  <head>
          <script>
                  location.replace(location.href.replace("https://","http://"));
          </script>
  </head>
  <body>
          <noscript><meta http-equiv="refresh" content="0;url=http://www.baidu.com/"></noscript>
  </body>
  </html>
*/
```