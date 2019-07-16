package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 分组路由
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")                     
	{
			v1.POST("/login", loginEndpoint)              //路由分组，访问地址变为：/v1/login
			v1.POST("/submit", submitEndpoint)
			v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
			v2.POST("/login", loginEndpoint)
			v2.POST("/submit", submitEndpoint)
			v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}

func loginEndpoint(){
	
}