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
