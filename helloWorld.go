package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-study/util"
)

func mid(c *gin.Context) {

	name := c.Query("name")
	if name == "cfl" {
		c.JSON(200, gin.H{
			"message": "拦截了",
		})
		c.Abort()
		return
	}
	fmt.Println("我还是继续了")
	c.Next()
}
func main() {
	fmt.Println(util.GetUName())
	fmt.Println("Hello 11World")
	r := gin.Default()
	r.Use(mid)
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("--进入了ping 方法---")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":5000")
}
