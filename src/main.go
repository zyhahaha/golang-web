package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// 重定向
func main() {
	r := gin.Default()
	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	// Listen and server on 0.0.0.0:8080
	r.Run(":8081")
}
