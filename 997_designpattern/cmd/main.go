package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/task1", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.JSON(200, gin.H{
			"message": "task1",
		})
	})

	r.GET("/task2", func(c *gin.Context) {
		time.Sleep(time.Second * 4)
		c.JSON(200, gin.H{
			"message": "task2",
		})
	})

	r.GET("/task3", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		c.JSON(200, gin.H{
			"message": "task3",
		})
	})

	r.GET("/task4", func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		c.JSON(200, gin.H{
			"message": "task4",
		})
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
