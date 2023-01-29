# What is XSS Filter

XSS's naming was defined to resolve confusing with CSS of HTML. Cause XSS is Cross Site Scripting, ( CSS is much more proper.). 
If we didn't set the xss filter on your api, a hacker or anyone can control your website easily.

> [Cross Site Scripting 에 대해서](https://crazykim2.tistory.com/694)

# The way how to use XSS filter in golang 

if we are using gin framework for web server, we can apply xss filter with below filter.

```go
package main

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// Applying Zap Logging
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	r := gin.Default()

	// include as standard middleware
	xssMdlwr := &xss.XssMw{
		FieldsToSkip: []string{"password", "create_date", "token"},
		BmPolicy:     "UGCPolicy",
	}
	r.Use(xssMdlwr.RemoveXss())

	r.GET("/ping", func(c *gin.Context) {

		value, ok := c.GetQuery("value")
		if ok {
			sugar.Infof("filtered xss value : %s", value)
		}

		token, ok := c.GetQuery("token")
		if ok {
			sugar.Infof("not filtered xss token : %s", token)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
```

```shell
[GIN] 2023/01/29 - 16:44:42 | 200 |       2.936ms |       127.0.0.1 | GET      "/ping?value=<>34234&token=<>"
{"level":"info","ts":1674978292.218557,"caller":"028_go_xss_filter/main.go:28","msg":"filtered xss value : &lt;&gt;34234"}
{"level":"info","ts":1674978292.218605,"caller":"028_go_xss_filter/main.go:33","msg":"not filtered xss token : <>12314"}
```

> [xss middleware](https://pkg.go.dev/github.com/dvwright/xss-mw)